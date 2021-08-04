package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Beginning")
	defer fmt.Println("Defered")
	panic("Something wrong happenned")

	/*fmt.Println("Termino")
	  The last Println is unreachable code as the panic stops entirely the work flow.
	  Recommend not using panic unless it's extremely important that the program stops
	  when certain task is not done properly.
	  Panic is always executed after defer.
	*/
}
