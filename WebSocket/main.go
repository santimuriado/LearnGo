/* First a standard http connection is built and then it's upgraded to a WebSocket connection. */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Reader which recieves sent messages to endpoint websocket.
func reader(conn *websocket.Conn) {
	for {
		//Read message.
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//Return true to keep it simple.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//Upgrade the connection to WebSocket.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected")

	err = ws.WriteMessage(1, []byte("Hello World"))
	if err != nil {
		log.Println(err)
	}

	//Listens for new messages.
	reader(ws)
}

func setupRoutes() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
