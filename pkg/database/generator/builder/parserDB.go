package builder

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type parseDb struct {
	Name   string
	Tables []string
}

type Parser struct {
	file    *ast.File
	db      parseDb
	imports map[string]string
}

func NewParser(fileName string) Parser {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	return Parser{
		file:    f,
		imports: map[string]string{},
		db:      parseDb{},
	}
}

func (p *Parser) Parse() (sqlTable SqlDatabase) {
	p.extractImports()
	p.getDatabase()
	sqlTable.Tables = p.parseTables()
	sqlTable.DbName = p.db.Name
	return
}

func (p *Parser) getDatabase() {
	ast.Inspect(p.file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.FuncDecl:
			p.db.Name = x.Name.Name
			p.parseTable(x)
			return false
		}
		return true
	})
}

func (p *Parser) parseTable(x ast.Node) {
	ast.Inspect(x, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.ReturnStmt:
			configs := x.Results[0].(*ast.CompositeLit).Elts
			keyValueExpr := configs[0].(*ast.KeyValueExpr)
			key := keyValueExpr.Key.(*ast.Ident).Name

			if key == "Tables" {
				value := keyValueExpr.Value
				elements := value.(*ast.CompositeLit).Elts

				for e := range elements {
					selector := elements[e].(*ast.CallExpr).Fun.(*ast.SelectorExpr)
					pkg := selector.X.(*ast.Ident).Name
					sel := selector.Sel.Name
					p.db.Tables = append(p.db.Tables, pkg+"."+sel)
				}

			}
			return false
		}
		return true
	})
}

func (p *Parser) extractImports() {
	ast.Inspect(p.file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.ImportSpec:
			path := strings.Replace(x.Path.Value, "\"", "", 2)
			parts := strings.Split(path, "/")
			ident := parts[len(parts)-1]
			p.imports[ident] = path
		}
		return true
	})
}

func (p *Parser) parseTables() []SqlTable {
	tables := make([]SqlTable, len(p.db.Tables))
	for index, ident := range p.db.Tables {
		parts := strings.Split(ident, ".")
		fileName := fmt.Sprintf("%s/src/%s/%s.go", os.Getenv("GOPATH"), p.imports[parts[0]], parts[1])
		tableParser := NewTableParser(fileName)
		tables[index] = tableParser.parse()
	}
	return tables
}
