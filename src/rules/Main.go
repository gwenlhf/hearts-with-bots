package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type GameMoveWrapper struct {
	Game Game
	Move Move
}


/*func webSocketHandler (w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	c.WriteMessage(websocket.TextMessage, []byte("it worked?"))
	return
}*/

func newGameHandler (w http.ResponseWriter, r *http.Request) {
	by, err := json.Marshal(NewGame())
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(by)
	return
}

func gameMoveHandler (w http.ResponseWriter, r *http.Request) {
	var rbody []byte
	q, err := r.Body.Read(rbody)
	if q == 0 || err != nil {
		log.Println(err)
		return
	}
	var objbod GameMoveWrapper
	err = json.Unmarshal(rbody, &objbod)
	if err != nil {
		log.Println(err)
		return
	}
	err = objbod.Game.Play(objbod.Move)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func main() {
	http.HandleFunc("/new", newGameHandler)
	http.HandleFunc("/mov", gameMoveHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("shit: ", err)
	}
}