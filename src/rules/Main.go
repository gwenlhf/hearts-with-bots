package main

import (
	"log"
	"flag"
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
)

type GameMoveWrapper struct {
	Game Game
	Move Move
}

type ErrorObject struct {}

// uncomment for training
// var trainiter = flag.Int("step", 989, "don't worry about it")

// newGameHandler is a REST endpoint, which returns the JSON Game result
// of calling NewGame(), and making the first move (the Two of Clubs).
func newGameHandler (w http.ResponseWriter, r *http.Request) {
	game := NewGame()
	game.Play(Move{
		game.ToPlay,
		Card{Clubs, Two},
	})
	by, err := json.Marshal(game)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(by)
	return
}

// gameMoveHandler is a REST endpoint, which takes a JSON GameMoveWrapper,
// and returns the JSON result of calling Game.Play(Move),
// or an ErrorObject if the Move is invalid.
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
	rand.Seed(time.Now().UnixNano())

	// for running games
	http.HandleFunc("/new", newGameHandler)
	http.HandleFunc("/mov", gameMoveHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("fatal: ", err)
	}

	// for creating training data
	/*flag.Parse()
	CreateTrainingData(*trainiter)*/
}