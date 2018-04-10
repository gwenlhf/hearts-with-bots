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

type ErrorObject struct {}


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
	var objbod GameMoveWrapper
	err := json.NewDecoder(r.Body).Decode(&objbod)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		by, _ := json.Marshal(new(ErrorObject))
		w.Write(by)
		return
	}
	err = objbod.Game.Play(objbod.Move)
	if err != nil {
		w.WriteHeader(400)
		by, _ := json.Marshal(new(ErrorObject))
		w.Write(by)
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	by, err := json.Marshal(objbod.Game)
	if err != nil {
		w.Write(by)
		w.WriteHeader(500)
		by, _ := json.Marshal(new(ErrorObject))
		w.Write(by)
		log.Println(err)
		return
	}
	w.Write(by)
	return
}

func main() {
	/*http.HandleFunc("/new", newGameHandler)
	http.HandleFunc("/mov", gameMoveHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("shit: ", err)
	}*/
	createTrainingData(1000)
}