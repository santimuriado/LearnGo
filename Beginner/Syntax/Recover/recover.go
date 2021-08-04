package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Beginning")
	panicker()
	fmt.Println("End")
}

func panicker() {

	fmt.Println("Panicking starts")

	//The recover function allows you to not stop the work flow with panic which would be
	//counterproductive in most cases.
	defer func() {
		if err := recover(); err != nil {
			//Logging error.
			log.Println("Error:", err)
		}
		//The parenthesis are necessary to call the function.
	}()
	panic("Something bad happenned")

	//The Println is unreachable code.
	//fmt.Println("No more panicking")

	//Only useful in defer functions.
}
