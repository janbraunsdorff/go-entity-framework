package main

import (
	"github.com/janbraunsdorff/demo/pkg/database/generator/builder"
	"os"
)

func main() {
	parser := builder.NewParser(os.Getenv("GOFILE"))
	db := parser.Parse()

	generator := builder.NewGenerator(db)
	generator.Generate()
}
