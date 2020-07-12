package entity

import "github.com/janbraunsdorff/demo/pkg/database/generator/item"

func Orders() (t item.Table) {
	return item.Table{
		Columns: []item.Field{
			{ // Columns in DB
				Name:      "OrderId", // name of column
				IsPrimary: true,      // mark as primary key
				DataType:  "int",     // datatype in DB (in entity -> int)
			},
			{
				Name:     "OrderNumber", // name of column
				DataType: "varchar(10)", // datatype in DB (in entity -> string)
				NotNull:  true,          // not allowed to be null
			},
			{
				Name:     "PersonId", // name of column
				DataType: "int",      // datatype in DB (in entity -> int)
				NotNull:  true,       // not allowed to be null
			},
		},
		Foreign: []item.ForeignKey{
			{ // referential integrity
				FieldName:        []string{"PersonId"}, // this fields
				ForeignTableName: "Person",             // foreign table
				ForeignFieldName: []string{"PersonId"}, // foreign fields
			},
		},
		Synthetic: []item.Table{ //Synthetic attribute (not in DB )
			Person(), //Single struct of type Person
		},
	}
}
