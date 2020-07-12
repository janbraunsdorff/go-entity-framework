package builder

import (
	. "github.com/janbraunsdorff/demo/pkg/test"
	"go/parser"
	"go/token"
	"testing"
)

func TestGetTableFromFile(t *testing.T) {
	fileName := "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/correctDbDefinition0Table.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	p := Parser{
		file:    f,
		imports: map[string]string{},
		db:      parseDb{},
	}

	p.getDatabase()

	Assert(t).That(p.db.Name).IsEqualTo("Test0DB", "database name is not correct")
	Assert(t).That(len(p.db.Tables)).IsEqualTo(1, "database should have one table")
	Assert(t).That(p.db.Tables[0]).IsEqualTo("entity.EmptyTable", "database should have a table with name 'Empty table'")
}

func TestCallParseTables(t *testing.T) {
	fileName := "/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/correctDbDefinition0Table.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	p := Parser{
		file:    f,
		imports: map[string]string{},
		db:      parseDb{},
	}

	p.Parse()

	Assert(t).That(p.db.Name).IsEqualTo("Test0DB", "database name is not correct")
	Assert(t).That(len(p.db.Tables)).IsEqualTo(1, "database should have one table")
	Assert(t).That(p.db.Tables[0]).IsEqualTo("entity.EmptyTable", "database should have a table with name 'Empty table'")
}
