package entity

import (
	"github.com/janbraunsdorff/demo/pkg/database/generator/item"
)

func Person() (t item.Table) {
	return item.Table{
		Columns: []item.Field{
			{ // Columns in DB
				Name:      "PersonId", // name of column
				IsPrimary: true,       // mark as primary key
				DataType:  "int",      // datatype in DB (in entity -> int)
			},
			{
				Name:     "LastName",    // name of column
				DataType: "varchar(60)", // datatype in DB (in entity -> string)
				NotNull:  true,          // not allowed to be null
			},
			{
				Name:     "FirstName",   // name of column
				DataType: "varchar(60)", // datatype in DB (in entity -> string)
				NotNull:  true,          // not allowed to be null
			},
			{
				Name:     "Age",       // name of column
				DataType: "int",       // datatype in DB (in entity -> int)
				Check:    "Age >= 18", // check clauses
			},
		},
		Synthetics: []item.Table{ //Synthetic attribute (not in DB )
			Orders(), //Slice of Person
		},
	}
}
