package main

import (
	"fmt"
	"errors"
	"math/rand"
)

type Suit int
const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

type Value int
const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Side int
const (
	North = iota
	East
	South
	West
)

type Card struct {
	Suit Suit
	Value Value
}

type Player struct {
	Side Side
	Hand map[Card]struct{}
	Points map[Card]struct{}
	Total int16
}

type Game struct {
	Players []Player
	Trick []Card
	HeartsBroken bool
	ToPlay Side
}

type Move struct {
	Side Side
	Card Card
}

func NewDeck () []Card {
	arr := new ([52]Card)
	for k, _ := range(arr) {
		arr[k] = Card{
			Suit(k/13), 
			Value(k % 13),
		}
	}
	rand.Shuffle(52, func (i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr[0:len(arr)-1]
}

func NewGame () Game {
	players := new ([4]Player)
	trick := new ([4]Card)
	deck := NewDeck()

	i, j := 0, 12
	for k, _ := range(players) {
		players[k].Hand = make(map[Card]struct{})
		for _, card := range(deck[i:j]) {
			players[k].Hand[card] = struct{}{}
		}
		players[k].Side = Side(k)
		players[k].Points = make(map[Card]struct{})
		players[k].Total = 0
		i += 13
		j += 13
	}
	game := Game{
		players[0:3],
		trick[0:0],
		false,
		StartingPlayer(players[0:3]),
	}
	game.Play(Move{
		game.ToPlay,
		Card{Clubs, Two},
	})
	return game
}

func StartingPlayer (players []Player) Side {
	for _, player := range(players) {
		for card, _ := range(player.Hand) {
			if card.Suit == Clubs && card.Value == Two {
				return player.Side
			}
		}	
	}
	panic ("No valid starting player")
}

func ( side Side ) Next () Side {
	if side == West {
		return North
	} else {
		return side + 1
	}
}

func ( game *Game ) Play (move Move) (err error) {
	if err, ok := game.ValidMove(move); !ok {
		return err
	}
	game.Trick = append(game.Trick, move.Card)
	delete(game.Players[move.Side].Hand, move.Card)

	if len(game.Trick) == 4 {
		game.Score()
	} else {
		game.ToPlay = game.ToPlay.Next()
	}

	return
}

// TODO : Side does not correlate with trick position
func ( game *Game ) Score () () {
	leader := game.ToPlay.Next()
	lead := game.Trick[0]
	// determine the winner of the trick
	for side, card := range(game.Trick) {
		if card.Suit == lead.Suit && card.Value > lead.Value {
			leader, lead = Side(leader + side), card
		}
	}
	// move the cards to that player's points
	points := game.Players[leader].Points
	for _, card := range(game.Trick) {
		points[card] = struct{}{}
	}
	// reset the trick
	game.Trick = game.Trick[0:0]
}

func ( game *Game ) ValidMove (move Move) (err error, ok bool) {
	// is it your turn?
	if move.Side != game.ToPlay {
		err = errors.New(fmt.Sprintf("Not %s's turn", move.Side))
		return err, false
	}

	player := game.Players[move.Side]

	// do you have the card?
	if _, ok := player.Hand[move.Card]; !ok {
		err = errors.New(fmt.Sprintf("%s does not have card %s", move.Side, move.Card))
		return err, false
	}

	// are you following the trick, if able?
	if len(game.Trick) > 0 && move.Card.Suit != game.Trick[0].Suit {
		for card, _ := range(player.Hand) {
			if card.Suit == game.Trick[0].Suit {
				err = errors.New(fmt.Sprintf("%s is able to follow trick", move.Side))
				return err, false
			}
		}
	}

	// if breaking hearts, are you legally able to?
	if !game.HeartsBroken && (move.Card.Suit == Hearts || move.Card == Card{Spades,Queen}) {
		for card, _ := range(player.Hand) {
			if card.Suit != Hearts && (card != Card{Spades,Queen}) {
				err = errors.New(fmt.Sprintf("%s cannot break hearts yet", move.Side))
				return err, false
			}
		}
	}
	// all good
	return nil, true
}