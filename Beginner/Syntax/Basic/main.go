package main

import (
	"fmt"
	"strconv"
	//"math"
)

//Iota uses incrementing numbers from 0.
const (
	errorIota = iota
	d
	e
	f
)

//The "_" ignores the first value.
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EZ
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNa
	canSeeSa
)

func main() {

	//VARIABLES
	name := "Santiago Muriado"
	floatNumber := 42.5
	intNumber := 23

	//Default value for variables is always 0.
	var testNumber int
	var testConv string

	fmt.Printf("%v, %T\n", name, name)

	fmt.Printf("%v, %T\n", floatNumber, floatNumber)

	fmt.Printf("%v, %T\n", testNumber, testNumber)

	floatNumber = float64(intNumber)
	fmt.Printf("%v, %T\n", floatNumber, floatNumber)

	testConv = strconv.Itoa(intNumber)
	fmt.Printf("%v,%T\n", testConv, testConv)

	//BOOLEANS

	//Initializes m as a boolean and compares it in the same line.
	m := 1 == 2
	fmt.Printf("%v,%T\n", m, m)

	//unsigned int
	var aNumber uint = 25
	fmt.Printf("%v,%T\n", aNumber, aNumber)

	//Operations in binary.
	a := 23
	b := 43
	fmt.Println(a & b)
	fmt.Println(a | b)

	c := 8
	fmt.Println(c << 3) // 8=2^3 => 2^3x2^3=64
	fmt.Println(c >> 3) // 2^3/2^3 = 1

	//COMPLEX NUMBERS.
	var complexNumber complex64 = 2 + 3i //Also complex128.
	fmt.Printf("%v, %T\n", complexNumber, complexNumber)
	//Real part of the complex number.
	fmt.Printf("%v, %T\n", real(complexNumber), real(complexNumber))
	//Imaginary part of the complex number.
	fmt.Printf("%v, %T\n", imag(complexNumber), imag(complexNumber))

	//Another declaration of a complex number.
	var complexNumber2 complex128 = complex(5, 12) //complexNumber2 = 5 + 2i.
	fmt.Printf("%v, %T\n", complexNumber2, complexNumber2)

	//STRINGS
	word := "a string"
	fmt.Printf("%v, %T\n", word, word)

	//Strings are aliases for bytes.
	//Returns the third char without counting spaces.
	fmt.Printf("%v, %T\n", string(word[3]), string(word[3]))

	//Concatenation.
	anotherWord := "another string"
	fmt.Printf("%v, %T\n", word+anotherWord, word+anotherWord)

	//You can divide them in bytes slices.
	sliceBytes := []byte(word)
	//Returns the corresponding ascii value of every char.
	fmt.Printf("%v, %T\n", sliceBytes, sliceBytes)

	//Lots of functions in golang work with byte slices and that makes them much more maleables than
	//in other languages.

	//Only accepts a single char just as a char in C++.
	//They are aliases for int32 and require special methods to read them as strings.Reader#ReadRune
	runeExample := 'e'
	fmt.Printf("%v,%T\n", runeExample, runeExample)

	//CONSTANTS

	//Can't assign a constant to a calculation because the calculation is made in runtime
	//and not in compile time.

	//When you have 2 constants with the same name there is a case of Shadowing.
	//The one that prevails depends on the scope we are working in.

	const myConstant = 16
	var myConstant2 int16 = 32
	//Casts myConstant into an int16.
	fmt.Printf("%v, %T\n", myConstant+myConstant2, myConstant+myConstant2)

	//Test IOTA and CONSTANTS.
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)
	//Works a lot like enum.

	testIota := d
	//Should be true.
	fmt.Printf("%v\n", testIota == 1)

	fileSize := 4000000000.
	//%.2f its a float number with 2 decimal spaces.
	fmt.Printf("%.2fGB \n", fileSize/GB)

	//This is called bit masking. This is very useful for storing data in a very efficient way.
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	//It compares the bits in isAdmin and in roles.
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Can see financials? %v\n", canSeeFinancials&roles == canSeeFinancials)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)

	//In Golang the while loop doesnt exist as it is another "shape" of a for loop.
	//However the for loop has multiple syntax.

	//USUAL
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//VARIABLE OUTSIDE THE LOOP.
	x := 5
	for ; x < 10; x++ {
		fmt.Println(x)
	}
	//WHILE LOOP BUT WITH A FOR.
	y := 10
	for y < 15 {
		fmt.Println(y)
		y++
	}
	//DO WHILE.
	z := 15
	for {
		fmt.Println(z)
		z++
		if z == 20 {
			break //KEYWORD
		}
	}
	//FOR CONTINUE VARIATION
	//Works pretty well for filtering.
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//LABELS
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop //It signals where you want to break the loop.
			}
		}
	}
	//FOR IN COLLECTIONS
	testSlice := []int{1, 2, 3}
	for k, v := range testSlice { //k of Key and v of Value.
		fmt.Println(k, v)
	}
	//Works the same way for arrays.

	//If i dont want the keys i can just use the blank identifier "_".
	//Example with blank identifier.
	testSlice2 := []int{1, 2, 3}
	for _, valor := range testSlice2 {
		fmt.Println(valor)
	}

	//Example without using values.
	testSlice3 := []int{1, 2, 3}
	for k := range testSlice3 {
		fmt.Println(k)
	}

}
