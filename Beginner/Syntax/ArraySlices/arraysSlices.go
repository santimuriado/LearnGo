package main

import (
	"fmt"
)

func main() {

	//Adyacent in memory.
	grades := [...]int{10, 8, 9, 5}
	fmt.Printf("Grades: %v\n", grades)

	var students [4]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Santiago"
	students[1] = "Lucio"
	students[2] = "Muriado"
	students[3] = "Alejandro"
	fmt.Printf("Students: %v\n", students[2])
	fmt.Printf("Student total:  %v\n", len(students)) //Len:cantidad de elementos en array

	//Golang copies the whole array but you can use pointers as in c++.
	array := [...]int{1, 2, 3}
	barray := &array
	barray[1] = 5
	fmt.Println(array)
	fmt.Println(barray)
	slice := []int{1, 2, 3, 4, 7, 3, 1, 0}
	fmt.Println(slice)
	fmt.Printf("Slice lenght: %v\n", len(slice))
	fmt.Printf("Slice capacity: %v\n", cap(slice))

	//Slices are passed as a reference.
	bslice := slice
	bslice[2] = 5
	fmt.Println(slice)
	fmt.Println(bslice)

	c := slice[:]   //slice of every element.
	d := slice[3:]  //Slice from the 4th element.
	e := slice[:6]  //Slice up to the 6th element included.
	f := slice[3:6] //Slice from the 4th element up to the 6th element included.

	//This changes the memory block so c,d,e and f are all changed.
	//Same as C++.
	f[2] = 38
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	//Specify the slice capacity.
	a := make([]int, 3, 100)

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := []int{}

	fmt.Println(b)
	fmt.Printf("b length: %v\n", len(b))
	fmt.Printf("b capacity: %v\n", cap(b))

	//Add elements to a slice.
	b = append(b, 1, 2, 3, 4, 5)

	fmt.Println(b)
	fmt.Printf("b length: %v\n", len(b))
	fmt.Printf("b capacity: %v\n", cap(b))

	//Concatenate 2 slices.
	b = append(b, []int{6, 7, 8}...)

	fmt.Println(b)
	fmt.Printf("b length: %v\n", len(b))
	fmt.Printf("b capacity: %v\n", cap(b))

	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(testSlice)

	//Removes the first element.
	g := testSlice[1:]
	//Removes the last element.
	h := testSlice[:len(testSlice)-1] //Retira el ultimo elemento

	fmt.Println(g)
	fmt.Println(h)

	//How to remove an element thats neither the first nor the last.

	/*
		b := append(a[:2], a[3:]...)
		fmt.Println(b)
	*/

}
