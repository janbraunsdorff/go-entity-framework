//go:generate go run ../generator

package database_description

import (
	"github.com/janbraunsdorff/demo/pkg/database/database-description/entity"
	"github.com/janbraunsdorff/demo/pkg/database/generator/item"
)

func Shop() (d item.Database) {
	return item.Database{ // database definition
		Tables: []item.Table{ // all tables
			entity.Person(), // add table person to DB
			entity.Orders(), // add table orders to DB
		},
	}
}
