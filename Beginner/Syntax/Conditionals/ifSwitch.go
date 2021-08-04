package main

import (
	"fmt"
)

func main() {

	number := 70
	otherNumber := 70

	//The "else" has to be in the same line as the closing bracket.
	if number < otherNumber {
		fmt.Println("other number is bigger")
	} else if number == otherNumber {
		fmt.Println("same number")
	} else {
		fmt.Println("number is bigger than otherNumber")
	}

	//Switch works the same way as in C++.
	//Cases can be stuck together. Ex: case 1,5,10:
	//												code

	i := 10

	// Logic comparisons can be made when tagless syntax is used.
	//Breaks are implicit.
	//Fallthrough doesnt respond to logic so the next case is going to be called either way.
	//Not really recommended or useful.
	switch {
	case i < 20:
		fmt.Println("Less than 20")
		fallthrough
	case i > 20:
		fmt.Println("Bigger than 20")
	default:
		fmt.Println("Either way")
	}

	//TYPE SWITCH.

	var f interface{} = 'a'
	switch f.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float")
	case [3]int:
		fmt.Println("array")
	default:
		fmt.Println("another type")

	}
}
