package main

import (
	"fmt"
	_ "github.com/lib/pq"
	generator "sqlingo-gen-postgres/generator"
)

func main() {
	generator.ParseArgs()
	code, err := generator.Generate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code)
}
