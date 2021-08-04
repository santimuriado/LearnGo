package main

import (
	"log"
	"net/http"
)

func main() {

	/*Serving from static folder. If the complexity of the server gets high enough, more than probable
	  the static files will be in another folder of their own. */
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//Normal way to open a server.
	//log.Fatal(http.ListenAndServe(":8080", nil))

	//Safe way to open a server. CRT and key have to be generated to access.
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))

}
