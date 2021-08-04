package main

import "fmt"

func calculate(number int) (result int) {
	result = number + 2
	return result
}

func main() {
	number := 3
	number = calculate(number)
	fmt.Println(number)
}

/*go test -cover tells you the porcentage coverage.
  go test -coverprofile=coverage.out builds an .out file with the coverage details.
  go tool cover -html=coverage.out opens the .out file in html format.
*/
