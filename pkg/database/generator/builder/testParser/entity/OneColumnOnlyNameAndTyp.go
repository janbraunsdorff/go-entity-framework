package entity

import "github.com/janbraunsdorff/demo/pkg/database/generator/item"

func OneColumnNameAndTyp() (i item.Table) {
	return item.Table{
		Columns: []item.Field{
			{
				Name:     "LastName",
				DataType: "varchar(10)",
			},
		},
		Foreign: []item.ForeignKey{
			// getPerson to order
			{
				FieldName:        []string{"PersonId"}, // this table
				ForeignTableName: "Person",
				ForeignFieldName: []string{"PersonId"},
			},
		},
	}
}
