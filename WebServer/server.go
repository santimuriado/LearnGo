package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.RWMutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
