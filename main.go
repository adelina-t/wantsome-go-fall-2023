package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message  string
	UserName string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var userConnections = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from my server!")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("got error upgrading connection %s\n", err)
		return
	}
	defer conn.Close()

	userConnections[conn] = true

	for {
		var msg Message = Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("got error reading message %s\n", err)
			delete(userConnections, conn)
			return
		}
		broadcast <- msg
	}
}

func handleMsg() {
	for {
		msg := <-broadcast

		for client := range userConnections {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("got error broadcating message to client %s", err)
				client.Close()
				delete(userConnections, client)
			}
		}
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", handleConnections)

	go handleMsg()

	fmt.Println("starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error starting server %s", err)
		os.Exit(1)
	}
}
