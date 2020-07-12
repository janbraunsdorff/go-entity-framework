package entity

import "github.com/janbraunsdorff/demo/pkg/database/generator/item"

func TwoColumnNameAndTyp() (i item.Table) {
	return item.Table{
		Columns: []item.Field{
			{
				Name:     "FirstName",
				DataType: "varchar(10)",
			},
			{
				Name:     "LastName",
				DataType: "varchar(10)",
			},
		},
	}
}
