package main

type money int      // it creates a new data type with name money, and it is different than int
type dollar = money // alias // dollar and money are same type

func main() {
	var m money = 100
	var x int = int(m) // we can't directly assign money values to int type, as it is a new different type
	var d dollar = m
	//var b byte // alias example in standard lib
	//time.Duration() // new type example in standard lib

	_, _, _ = m, x, d // underscore would mark the vars as used
}
