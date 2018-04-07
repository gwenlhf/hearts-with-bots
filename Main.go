package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	game := NewGame()
	fmt.Print(json.Marshal(game))
/*	card := Card{Clubs,Two}
	fmt.Print(json.Marshal(card))*/
}