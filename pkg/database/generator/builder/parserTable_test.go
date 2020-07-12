package builder

import (
	. "github.com/janbraunsdorff/demo/pkg/test"
	"testing"
)

func TestParseEmptyTable(t *testing.T) {
	p := NewTableParser("/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/entity/EmptyTable.go")
	table := p.parse()

	Assert(t).That(table.TableName).IsEqualTo("EmptyTable", "Table should be the name 'EmptyTable'")
	Assert(t).That(len(table.Columns)).IsEqualTo(0, "Table has no columns")
}

func TestParseOneColumnNameAndType(t *testing.T) {
	p := NewTableParser("/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/entity/OneColumnOnlyNameAndTyp.go")
	table := p.parse()

	Assert(t).That(table.TableName).IsEqualTo("OneColumnNameAndTyp", "Table should be the name 'OneColumnNameAndTyp'")
	Assert(t).That(len(table.Columns)).IsEqualTo(2, "Table should be the name 'Name'")
	Assert(t).That(table.Columns[0]).IsEqualTo("LastName varchar(10)", "Table should be the name 'Name'")
}

func TestParseTwoColumnNameAndType(t *testing.T) {
	p := NewTableParser("/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/entity/TwoColumnOnlyNameAndTyp.go")
	table := p.parse()

	Assert(t).That(table.TableName).IsEqualTo("TwoColumnNameAndTyp", "Table should be the name 'TwoColumnNameAndTyp'")
	Assert(t).That(len(table.Columns)).IsEqualTo(2, "Table should be the name 'Name'")
	Assert(t).That(table.Columns[0]).IsEqualTo("FirstName varchar(10)", "Table should be the name 'Name'")
	Assert(t).That(table.Columns[1]).IsEqualTo("LastName varchar(10)", "Table should be the name 'Name'")
}

func TestParseFullDefinitionColumn(t *testing.T) {
	p := NewTableParser("/Users/janbraunsdorff/go/src/github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/entity/OneColumnWithFullDefinition.go")
	table := p.parse()

	Assert(t).That(table.TableName).IsEqualTo("FullDefinition", "Table should be the name 'FullDefinition'")
	Assert(t).That(len(table.Columns)).IsEqualTo(1, "Table should be the name 'Name'")
	Assert(t).That(table.Columns[0]).IsEqualTo("Name varchar(10) NOT NULL PRIMARY KEY UNIQUE AUTO_INCREMENT DEFAULT 'Jan' CHECK (Name >= 18)", "Table should be the name 'Name'")
}
