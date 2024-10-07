// only main could be a binary package,
// binary package, a package we can run

package main

import (
	"fmt"
	"github.com/username/reponame/sum"
) // modulename/packageName// to import a custom package

//https://google.github.io/styleguide/go/

//go run . or go run *.go // to run all the main package files in current dir

func main() {
	sum.Addition(3, 2)
	//sum.Total = 999 // avoid exporting global var, someone can change it in between
	fmt.Println(sum.Total)
	startup()

}
