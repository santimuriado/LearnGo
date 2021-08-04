package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//Defer runs at the end of the program.
	fmt.Println("Beginning")
	defer fmt.Println("Half")
	fmt.Println("End")

	//Defer runs in LIFO.

	//EXAMPLE.
	res, err := http.Get("http://www.google.com/robots.txt")

	//Error checking.
	if err != nil {
		//Log the error in case it exists.
		log.Fatal(err)
	}

	// It is a good practice to close the file with a defer after opening it so you dont forget to
	// close it.
	//However be careful opening multiple files as it uses resources.
	//Better to close files as they stop being used.
	defer res.Body.Close()

	//Parses and casts it into a byte string.
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
