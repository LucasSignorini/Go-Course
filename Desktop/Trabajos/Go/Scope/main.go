package main

import (
	"fmt"
	"myapp2/packageone"
)

var myVar = 1

func main() {
	var blockVar = 2
	packageone.PrintMe(myVar, blockVar)
	fmt.Println("Fin")
}
