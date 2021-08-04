package main

import (
	"fmt"
)

func main() {
	firstWord := "Hi"
	secondWord := "Everyone"

	sayHi(&firstWord, &secondWord)
	fmt.Println(secondWord)

	numberSum := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum of the numbers is: ", numberSum)

	number, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(number)

	//Anonymous function
	func() {
		fmt.Println("")
	}()

	//It calls the function at the same time as it is made.
	var f func() = func() {
		fmt.Println("Hello")
	}
	f()

	//Method.
	greet := greeter{
		greet: "Hello",
		name:  "Santiago",
	}
	greet.sayHiMethod()

}

type greeter struct {
	greet string
	name  string
}

//Method. Basically a function that only works in a specific scope, in this case a struct.
func (s greeter) sayHiMethod() {
	fmt.Println(s.greet, s.name)
}

func sayHi(firstWord, secondWord *string) {

	fmt.Println(*firstWord, *secondWord)
	*secondWord = "Santiago"
	fmt.Println(*firstWord, *secondWord)
}

//Variadic function. It takes an arbitrary number of ints in this case.
func sum(valores ...int) (resultado int) {
	fmt.Println(valores)
	for _, v := range valores {
		resultado += v
	}
	return
}

/*func sum(valores ...int) *int {
	fmt.Println(valores)
	resultado := 0
	for _, v := range valores {
		resultado += v
	}
	return &resultado

	Normally when you work in with variables in functions, they are stored in the execution stack but when the function ends
	the execution stack is erased to free up memory. Go recognizes the return of the pointer to something that was previously
	stored in the local stack and promotes the variable to be in the heap(shared memory). We can always return the address
	of the return value and Go does the rest.

}*/

func divide(a, b float64) (float64, error) {
	//It's recommended not to use panic becaause it stops the work flow.
	if b == 0.0 {
		return 0.0, fmt.Errorf("no se puede divide por 0")
	}
	return a / b, nil
}

// It's always better to use pointers as they are almost always smaller than the actual data they
// refer to.
