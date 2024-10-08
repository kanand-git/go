package main

import "fmt"

func main() {
	//map[key]value
	dictionary := make(map[string]string)

	dictionary["up"] = "above"
	dictionary["below"] = "down"

	fmt.Println(dictionary["up"])
	// maps in go are not ordered, don't rely on order when ranging over it
	for key, value := range dictionary {
		fmt.Println(key, value)
	}
	//Passing a map to the len function tells you the number
	//of key-value pairs in a map.
	fmt.Println(len(dictionary))
	delete(dictionary, "up")

	// access the value of key up
	v, ok := dictionary["up"]
	if !ok {
		fmt.Println("key up not found")
		return
	}
	fmt.Println(v)

	fmt.Println(dictionary)
}
