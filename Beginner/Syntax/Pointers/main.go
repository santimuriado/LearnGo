package main

import "fmt"

func main() {

	//Pointers work almost the same as in C++.

	var a int = 42
	//Points to the a address.
	var b *int = &a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, *b)
	//Modify a, modifying the b pointer.
	*b = 40
	fmt.Println(a, *b)

	//Adyacent memory blocks.
	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)
	//There is no such thing as pointer arithmetics in Go as in C++.

	//Pointers are initialized in nil.
	var ms *myStruct
	fmt.Println(ms)

	ms = new(myStruct)
	ms.foo = 42
	//Go does the deferencing without having to manipulate the pointer.
	fmt.Println(ms.foo)

	//Arrays are passed in value but slices and maps from reference.
	f := map[string]string{"word": "something", "table": "chair"}
	g := f
	fmt.Println(f, g)
	//It modifies both maps as intended.
	f["word"] = "something"
	fmt.Println(f, g)
}

type myStruct struct {
	foo int
}
