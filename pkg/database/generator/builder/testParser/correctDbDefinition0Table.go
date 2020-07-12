package testParser

import (
	"github.com/janbraunsdorff/demo/pkg/database/generator/builder/testParser/entity"
	"github.com/janbraunsdorff/demo/pkg/database/generator/item"
)

func Test0DB() (d item.Database) {
	return item.Database{
		Tables: []item.Table{
			entity.EmptyTable(),
		},
	}
}
