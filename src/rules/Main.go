package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	game := NewGame()
	fmt.Print(json.Marshal(game))
}