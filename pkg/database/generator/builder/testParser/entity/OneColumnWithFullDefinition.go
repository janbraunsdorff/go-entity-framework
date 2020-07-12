package entity

import "github.com/janbraunsdorff/demo/pkg/database/generator/item"

func FullDefinition() (i item.Table) {
	return item.Table{
		Columns: []item.Field{
			{
				Name:      "Name",
				AutoInc:   true,
				NotNull:   true,
				IsUnique:  true,
				IsPrimary: true,
				DataType:  "varchar(10)",
				Check:     "Name >= 18",
				DefValue:  "Jan",
			},
		},
	}
}
