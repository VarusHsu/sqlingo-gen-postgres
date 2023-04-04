package main

import (
	"fmt"
	"github.com/VarusHsu/sqlingo-gen-postgres/generator"
	_ "github.com/lib/pq"
)

func main() {
	generator.ParseArgs()
	code, err := generator.Generate()
	if err != nil {
		fmt.Println(err)
	
	}
	fmt.Println(code)
}
