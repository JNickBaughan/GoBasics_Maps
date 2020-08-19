package main

import "fmt"

func main() {
	// map with [string] for the keys and string for the values
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(numbers)
	// use the delete function to remove values
	delete(numbers, "three")
	fmt.Println(numbers)
	// declare an empty map
	var colors map[string]string
	fmt.Println(colors)
	// you cannot add to a map initialized like above
	// you will get the error "panic: assignment to entry in nil map"
	// use the make method below
	// what the point of this syntax
	// colors["red"] = "#FF0000"
	// colors["blue"] = "#0000FF"
	// fmt.Println(colors)

	// decalre an empty map with built in make function
	shapes := make(map[string]string)
	fmt.Println("shapes: ", shapes)
	// you must use the [] syntax to access the value
	// cannot use .{keyName} syntax like with a struct
	// this is because the keys aren't always string
	shapes["circle"] = "circle"
	fmt.Println("shapes: ", shapes)

	numberAsKey := make(map[int]string)
	// use [] syntax to access values
	// always use the [] syntax because numberAsKey.1 wouldn't work
	numberAsKey[1] = "one"
	fmt.Println("numberAsKey: ", numberAsKey)

}
