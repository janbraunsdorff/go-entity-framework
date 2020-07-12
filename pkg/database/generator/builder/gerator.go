package builder

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

type SqlDatabase struct {
	DbName string
	Tables []SqlTable
	Date   time.Time
}

type SqlTable struct {
	TableName  string
	Columns    []string
	ForeignKey []string
	Synthetic  []string
	Synthetics []string
}

type SqlConfig struct {
	Date time.Time
}

type GoEntity struct {
	Name       string
	Attributes []string
	Date       time.Time
}

type Generator struct {
	db                 SqlDatabase
	configTemplateIn   string
	configTemplateOut  string
	managerTemplateIn  string
	managerTemplateOut string
	entityPath         string
	entityIn           string
}

func NewGenerator(database SqlDatabase) Generator {
	return Generator{
		db:                database,
		configTemplateIn:  "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/template/dbConfig.tmpl",
		configTemplateOut: "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/access/config.go",

		managerTemplateIn:  "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/template/manager.tmpl",
		managerTemplateOut: "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/access/manager.go",

		entityIn:   "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/template/entity.tmpl",
		entityPath: "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/access/entity",
	}
}

func (g *Generator) Generate() {
	fmt.Printf("1. %s", "remove old files ")
	g.clean()
	fmt.Printf("%s\n", "done")

	fmt.Printf("2. %s", "crate config ")
	g.createConfig()
	fmt.Printf("%s\n", "done")

	fmt.Printf("3. %s", "crate manager ")
	g.createManager()
	fmt.Printf("%s\n", "done")

	fmt.Printf("4. %s", "crate entities ")
	g.createEntities()
	fmt.Printf("%s\n", "done")
}

func (g *Generator) createConfig() {
	file, err := ioutil.ReadFile(g.configTemplateIn)
	if err != nil {
		panic(err)
	}

	parse, err := template.New("Config").Parse(string(file))
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	err = parse.Execute(&buffer,
		SqlConfig{
			Date: time.Now(),
		})
	if err != nil {
		panic(err)
	}

	g.saveFile(buffer, g.configTemplateOut)
}

func (g *Generator) createManager() {
	for t := range g.db.Tables {
		for c := range g.db.Tables[t].Columns {
			if c <= len(g.db.Tables[t].Columns)-2 {
				g.db.Tables[t].Columns[c] = g.db.Tables[t].Columns[c] + ","
			}
			g.db.Tables[t].Columns[c] = g.db.Tables[t].Columns[c] + "\n\t\t   "
		}
	}

	g.db.Date = time.Now()

	file, err := ioutil.ReadFile(g.managerTemplateIn)
	if err != nil {
		panic(err)
	}

	parse, err := template.New("Config").Parse(string(file))
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	err = parse.Execute(&buffer, g.db)
	if err != nil {
		panic(err)
	}

	g.saveFile(buffer, g.managerTemplateOut)
}

func (g *Generator) saveFile(buffer bytes.Buffer, path string) {
	p, err := format.Source(buffer.Bytes())
	if err != nil {
		// handle error
	}

	err = ioutil.WriteFile(path, p, 0644)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) createEntities() {
	for i := range g.db.Tables {
		g.createEntity(g.db.Tables[i])
	}
}

func (g *Generator) clean() {
	g.delete(g.configTemplateOut)
	g.delete(g.managerTemplateOut)
	g.deleteDir(g.entityPath)
}

func (g *Generator) delete(file string) {
	if _, err := os.Stat(file); os.IsExist(err) {
		err := os.Remove(file)
		if err != nil {
			panic(err)
		}
	}
}

func (g *Generator) deleteDir(file string) {
	_ = os.RemoveAll(file)
	_ = os.Mkdir(file, 0777)
}

func (g *Generator) createEntity(table SqlTable) {
	entity := GoEntity{
		Date:       time.Now(),
		Name:       table.TableName,
		Attributes: []string{},
	}

	for i := range table.Columns {
		parts := strings.Split(table.Columns[i], " ")
		cName := parts[0]
		cType := strings.ToLower(parts[1])
		goType := ""

		if strings.HasPrefix(cType, "char") {
			goType = "string"
		}
		if strings.HasPrefix(cType, "varchar") {
			goType = "string"
		}
		if strings.HasPrefix(cType, "tinytext") {
			goType = "string"
		}
		if strings.HasPrefix(cType, "text") {
			goType = "string"
		}
		if strings.HasPrefix(cType, "longtext") {
			goType = "string"
		}
		if strings.HasPrefix(cType, "mediumtext") {
			goType = "string"
		}
		if strings.HasPrefix(cType, "binary") {
			goType = "[]byte"
		}
		if strings.HasPrefix(cType, "varbinary") {
			goType = "[]byte"
		}
		if strings.HasPrefix(cType, "tinyblob") {
			goType = "[]byte"
		}
		if strings.HasPrefix(cType, "mediumblob") {
			goType = "[]byte"
		}
		if strings.HasPrefix(cType, "longblob") {
			goType = "[]byte"
		}

		if strings.HasPrefix(cType, "bit") {
			goType = "int"
		}
		if strings.HasPrefix(cType, "tinyint") {
			goType = "int"
		}
		if strings.HasPrefix(cType, "smallint") {
			goType = "int"
		}
		if strings.HasPrefix(cType, "mediumint") {
			goType = "int"
		}
		if strings.HasPrefix(cType, "int") {
			goType = "int"
		}

		if strings.HasPrefix(cType, "float") {
			goType = "float64"
		}
		if strings.HasPrefix(cType, "double") {
			goType = "float64"
		}
		if strings.HasPrefix(cType, "decimal") {
			goType = "float64"
		}
		if strings.HasPrefix(cType, "dec") {
			goType = "float64"
		}
		if strings.HasPrefix(cType, "bigint") {
			goType = "int"
		}
		if strings.HasPrefix(cType, "bool") {
			goType = "bool"
		}
		if strings.HasPrefix(cType, "boolean") {
			goType = "bool"
		}

		if strings.HasPrefix(cType, "date") {
			goType = "*time.time"
		}
		if strings.HasPrefix(cType, "time") {
			goType = "*time.time"
		}
		if strings.HasPrefix(cType, "year") {
			goType = "int"
		}

		entity.Attributes = append(entity.Attributes, cName+" "+goType)
	}

	entity.Attributes = append(entity.Attributes, table.Synthetics...)
	entity.Attributes = append(entity.Attributes, table.Synthetic...)

	file, err := ioutil.ReadFile(g.entityIn)
	if err != nil {
		panic(err)
	}

	parse, err := template.New("Entity").Parse(string(file))
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	err = parse.Execute(&buffer, entity)
	if err != nil {
		panic(err)
	}

	g.saveFile(buffer, g.entityPath+"/"+entity.Name+".go")

}
