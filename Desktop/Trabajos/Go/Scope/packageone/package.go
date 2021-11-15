package packageone

import "fmt"

var PackageVar int

func PrintMe(myVar, blockVar int) {
	fmt.Println(myVar, blockVar, PackageVar)
}
