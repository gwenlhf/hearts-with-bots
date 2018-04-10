package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func webSocketHandler (w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	c.WriteMessage(websocket.TextMessage, []byte("it worked?"))
	return
}

func main() {
	game := NewGame()
	fmt.Print(json.Marshal(game))/*
	http.HandleFunc("/ws", webSocketHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("shit: ", err)
	}*/
}