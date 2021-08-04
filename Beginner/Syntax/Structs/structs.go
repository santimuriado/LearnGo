package main

import (
	"fmt"
	"reflect"
)

//If variables or structs are to be used globally they have to be capitalized.
type School struct {
	name           string
	numberStudents int
	worldRanking   int
	studentsNames  []string
}

type Animal struct {
	name   string `max:"100"`
	origin string
}

type Pajaro struct {
	Animal
	speed  int
	canFly bool
}

func main() {

	aSchool := School{
		name:           "Brendans",
		numberStudents: 45,
		worldRanking:   1,
		studentsNames: []string{
			"Santiago Muriado",
			"Augusto Bottelli",
			"Augusto Montaldo",
			"Lorenzo Travaglini"},
	}

	aSchool.numberStudents = 72
	aSchool.studentsNames[1] = "Federico Caffaro"

	fmt.Println(aSchool)
	fmt.Println(aSchool.studentsNames)

	//Anonymous declaration of a struct.
	doctor := struct{ name string }{name: "Jony Gomez"}
	fmt.Printf("%v\n", doctor.name)

	// Structs are passed as value as oppossed to reference. They can be passed as reference as in
	// C++ with &.
	anotherDoctor := doctor
	anotherDoctor.name = "Tony Lopez"
	fmt.Printf("%v\n", anotherDoctor.name)

	//EMBEDDING
	//Go isn't OOP so it doesn't have inheritance, it has composition.

	testBird := Pajaro{}

	testBird.name = "Hawk"
	testBird.origin = "Mountain"
	testBird.speed = 40
	testBird.canFly = false
	fmt.Println(testBird)
	//To simulate behaviour interfaces are much better. Embedding is better for basic attributes.

	//TAGS
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

}
