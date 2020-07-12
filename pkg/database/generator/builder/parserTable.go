package builder

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"text/template"
)

type TableParser struct {
	file *ast.File
}

type columnDefinition struct {
	Name          string
	Type          string
	AutoIncrement string
	Check         string
	NotNull       string
	PrimaryKey    string
	Unique        string
	Default       string
}

type foreignKey struct {
	Fields    string
	Table     string
	RefFields string
}

func NewTableParser(fileName string) TableParser {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	return TableParser{file: f}
}

func (p *TableParser) parse() (table SqlTable) {
	entityDefinition := p.searchForEntity()
	table.TableName = entityDefinition.Name.Name
	restStmt := p.findReturnValue(entityDefinition.Body.List)
	table.Columns, table.ForeignKey, table.Synthetics = p.parseColumns(restStmt)

	return table
}

func (p *TableParser) searchForEntity() (fd *ast.FuncDecl) {
	ast.Inspect(p.file, func(node ast.Node) bool {
		if n, ok := node.(*ast.FuncDecl); ok {
			fd = n
			return false
		}
		return true
	})
	return
}

func (p *TableParser) findReturnValue(list []ast.Stmt) (retStmt *ast.ReturnStmt) {
	for i := range list {
		if w, ok := list[i].(*ast.ReturnStmt); ok {
			retStmt = w
			break
		}
	}
	return
}

func (p *TableParser) parseColumns(stmt *ast.ReturnStmt) (columns, foreignKeys, synthetics []string) {
	returnList := stmt.Results[0].(*ast.CompositeLit)

	if len(returnList.Elts) <= 0 {
		return
	}

	actions := map[string]func(val *ast.KeyValueExpr){
		"Columns": func(val *ast.KeyValueExpr) {
			p.checkIfIdentAndNotNil(val)

			for _, column := range val.Value.(*ast.CompositeLit).Elts {
				columns = append(columns, p.parseColumn(column))
			}
		},
		"Foreign": func(val *ast.KeyValueExpr) {
			p.checkIfIdentAndNotNil(val)
			for _, foreign := range val.Value.(*ast.CompositeLit).Elts {
				foreignKeys = append(foreignKeys, p.parseForeignKey(foreign))
			}
		},
		"Synthetics": func(val *ast.KeyValueExpr) {
			p.checkIfIdentAndNotNil(val)
			for i := range val.Value.(*ast.CompositeLit).Elts {
				name := fmt.Sprintf("%v", val.Value.(*ast.CompositeLit).Elts[i].(*ast.CallExpr).Fun.(*ast.Ident))
				synthetics = append(synthetics, fmt.Sprintf("%s []%s", name, name))
			}
		},
		"Synthetic": func(val *ast.KeyValueExpr) {
			p.checkIfIdentAndNotNil(val)
			for i := range val.Value.(*ast.CompositeLit).Elts {
				name := fmt.Sprintf("%v", val.Value.(*ast.CompositeLit).Elts[i].(*ast.CallExpr).Fun.(*ast.Ident))
				synthetics = append(synthetics, fmt.Sprintf("%s %s", name, name))
			}

		},
	}

	for i := range returnList.Elts {
		column := returnList.Elts[i].(*ast.KeyValueExpr)
		key := column.Key.(*ast.Ident).Name
		actions[key](column)
	}

	return
}

func (p *TableParser) checkIfIdentAndNotNil(val *ast.KeyValueExpr) {
	if w, ok := val.Value.(*ast.Ident); ok && w.Name == "nil" {
		panic("no columns are present")
	}
}

func (p *TableParser) parseColumn(column ast.Expr) string {
	cd := columnDefinition{}

	actions := map[string]func(val interface{}){
		"Name": func(val interface{}) {
			cd.Name = val.(string)
		},
		"DataType": func(val interface{}) {
			cd.Type = val.(string) + " "
		},
		"NotNull": func(val interface{}) {
			if p.checkIfTrue(val) {
				cd.NotNull = "NOT NULL "
			} else {
				cd.NotNull = ""
			}
		},
		"IsUnique": func(val interface{}) {
			if p.checkIfTrue(val) {
				cd.Unique = "UNIQUE "
			} else {
				cd.Unique = ""
			}
		},
		"IsPrimary": func(val interface{}) {
			if p.checkIfTrue(val) {
				cd.PrimaryKey = "PRIMARY KEY "
			} else {
				cd.PrimaryKey = ""
			}
		},
		"Check": func(val interface{}) {
			if p.checkIfPresent(val) {
				cd.Check = "CHECK (" + val.(string) + ") "
			} else {
				cd.Check = ""
			}
		},
		"DefValue": func(val interface{}) {
			if p.checkIfPresent(val){
				cd.Default = "DEFAULT '" + val.(string) + "' "
			} else {
				cd.Default = ""
			}
		},
	}

	for _, value := range column.(*ast.CompositeLit).Elts {
		config := value.(*ast.KeyValueExpr)
		name := config.Key.(*ast.Ident).Name
		var val interface{}

		if w, ok := config.Value.(*ast.BasicLit); ok {
			val = p.cleanBasicLiteral(w.Value)
		}

		if w, ok := config.Value.(*ast.Ident); ok {
			val = p.cleanBasicLiteral(w.Name)
		}

		actions[name](val)
	}

	return p.executeColumnTemplate(cd)
}

func (p *TableParser) checkIfPresent(val interface{}) bool {
	return val != nil && val.(string) != ""
}

func (p *TableParser) checkIfTrue(val interface{}) bool {
	return val.(string) == "true"
}

func (p *TableParser) executeColumnTemplate(cd columnDefinition) string {
	t, err := template.New("column").Parse(`
		{{.Name}} 
		{{.Type}}
		{{.PrimaryKey}}
		{{.NotNull}}
		{{.Unique}}
		{{.Default}}
		{{.Check}}
	`)

	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, cd)

	if err != nil {
		panic(err)
	}

	s := strings.ReplaceAll(tpl.String(), "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s[:len(s)-1]
}

func (p *TableParser) parseForeignKey(foreign ast.Expr) string {
	fk := foreignKey{
		Fields:    "",
		Table:     "",
		RefFields: "",
	}
	actions := map[string]func(*ast.KeyValueExpr){
		"FieldName": func(node *ast.KeyValueExpr) {
			columns := p.readFields(node)
			fk.Fields = strings.Join(columns, ",")
		},
		"ForeignFieldName": func(node *ast.KeyValueExpr) {
			columns := p.readFields(node)
			fk.RefFields = strings.Join(columns, ",")
		},
		"ForeignTableName": func(node *ast.KeyValueExpr) {
			fk.Table = strings.ReplaceAll(node.Value.(*ast.BasicLit).Value, "\"", "")
		},
	}

	for _, value := range foreign.(*ast.CompositeLit).Elts {
		config := value.(*ast.KeyValueExpr)
		name := config.Key.(*ast.Ident).Name

		actions[name](config)

	}

	return p.executeForeignKeyTemplate(fk)
}
func (p *TableParser) executeForeignKeyTemplate(fk foreignKey) string {
	t, err := template.New("column").Parse(`
		FOREIGN KEY ({{.Fields}}) REFERENCES {{.Table}}({{.RefFields}})
	`)

	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, fk)

	if err != nil {
		panic(err)
	}

	s := strings.ReplaceAll(tpl.String(), "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func (p *TableParser) readFields(node *ast.KeyValueExpr) []string {
	val := node.Value.(*ast.CompositeLit)
	var columns []string
	for i := range val.Elts {
		columnName := p.cleanBasicLiteral(val.Elts[i].(*ast.BasicLit).Value)
		columns = append(columns, columnName)
	}
	return columns
}

func (p *TableParser) cleanBasicLiteral(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}
