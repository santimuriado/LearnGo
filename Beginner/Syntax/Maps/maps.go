package main

import (
	"fmt"
)

func main() {

	//This is how maps are created without the data input.
	//cityPopulation := make(map[string]int)

	cityPopulation := map[string]int{
		"Buenos Aires": 4000000,
		"Nueva York":   35000000,
		"Paris":        45000000,
		"Londres":      20000000,
		"Berlin":       15000000,
		"Zurich":       70000000,
	}

	fmt.Println(cityPopulation)

	//It is possible to create maps from arrays but not from slices.
	m := map[[3]int]string{}
	fmt.Println(m)

	fmt.Println(cityPopulation["Paris"])

	//Append to the map.
	cityPopulation["Tokyo"] = 100000000
	fmt.Println(cityPopulation["Tokyo"])
	fmt.Println(cityPopulation)
	//Maps return value is not guaranteed, the return order is random.

	//Delete from map.
	delete(cityPopulation, "Zurich")
	fmt.Println(cityPopulation)
	//If the value searched doesn't exist the map returns a 0 but no error.

	//"Buenos Aires" typed incorrectly on purpose expecting a 0 and a false bool.
	population, ok := cityPopulation["Buenos Ares"]
	//This works to check if the element is included in the map.
	fmt.Println(population, ok)

	//Only interested in the boolean.
	_, ok = cityPopulation["Tokyo"]
	fmt.Println(ok)

	//Map size.
	fmt.Println(len(cityPopulation))

	//Maps are passed as reference.
	anotherMap := cityPopulation
	//This also deletes it from cityPopulation.
	delete(anotherMap, "Berlin")
	fmt.Println(cityPopulation)

	if testPopulation, ok := cityPopulation["Londres"]; ok {
		fmt.Println(testPopulation)
	}

	/*Logical operators are the same as in C++.
	  Go doesn't continue to evaluate logical comparisons when the conditions have already been met.
	  For example in a && comparison if the first comparison fails this causes what's called
	  a short circuit and stops evaluating.
	*/

}
