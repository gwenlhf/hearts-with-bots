// Hearts.go handles all of the game logic for modelling and playing the card game Hearts.
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
	Hand []Card
	Points []Card
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

// NewDeck returns a pseudorandom list of 52 Cards, one of each Suit/Value.
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

// NewGame returns a new Game, with pseudorandomized Cards.
func NewGame () Game {
	players := new ([4]Player)
	trick := new ([4]Card)
	deck := NewDeck()

	i, j := 0, 13
	for k, _ := range(players) {
		players[k] = Player{
			Side(k),
			deck[i:j],
			(new ([52]Card))[0:0],
			0,
		}
		i += 13
		j += 13
	}
	game := Game{
		players[0:4],
		trick[0:0],
		false,
		StartingPlayer(players[0:4]),
	}
	return game
}

// StartingPlayer returns the Side of the Player who has the Two of Clubs.
// StartingPlayer panics if no Player has the Two of Clubs.
func StartingPlayer (players []Player) Side {
	for _, player := range(players) {
		for _, card := range(player.Hand) {
			if card.Suit == Clubs && card.Value == Two {
				return player.Side
			}
		}	
	}
	panic ("No valid starting player")
}

func ( side Side ) Next ( n int ) Side {
	return Side(((int(side) + n) % 4))
}

// Play applies a Move to a Game.
// If the Move is invalid for the Game, Play returns an error and does nothing.
func ( game *Game ) Play (move Move) (err error) {
	if err, ok := game.ValidMove(move); !ok {
		return err
	}
	game.Trick = append(game.Trick, move.Card)

	// find and remove card from hand
	hand := game.Players[move.Side].Hand
	for i, card := range(hand) {
		if card == move.Card {
			hand[i] = hand[len(hand)-1]
			// probably a cleaner way to do this with pointers
			game.Players[move.Side].Hand = hand[:len(hand)-1]
			break
		}
	}

	if len(game.Trick) == 4 {
		game.Score()
	} else {
		game.ToPlay = game.ToPlay.Next(1)
	}

	return
}

// Score updates a Game with a complete trick, scoring the cards in that trick.
func ( game *Game ) Score () () {
	leader := game.ToPlay.Next(1)
	lead := game.Trick[0]
	// determine the winner of the trick
	for i, card := range(game.Trick) {
		if card.Suit == lead.Suit && card.Value > lead.Value {
			leader, lead = leader.Next(i), card
		}
	}
	// move the cards to that player's points
	leadplayer := game.Players[leader]
	for _, card := range(game.Trick) {
		leadplayer.Points = append(leadplayer.Points, card)
		if card.Suit == Hearts {
			leadplayer.Total += 1
		} else if (card == Card{Spades,Queen}) {
			leadplayer.Total += 13
		}
	}
	// reset the trick
	game.Trick = game.Trick[0:0]
	// handle moonshots
	if leadplayer.Total == 26 {
		leadplayer.Total = -26
	}
	// set next player to winner
	game.ToPlay = leadplayer.Side
}

// ValidMove determines if a given move is valid for some Game.
func ( game *Game ) ValidMove (move Move) (err error, ok bool) {
	// is it your turn?
	if move.Side != game.ToPlay {
		err = errors.New(fmt.Sprintf("Not %s's turn", move.Side))
		return err, false
	}

	player := game.Players[move.Side]

	// do you have the card?
	hasCard := false
	for _, card := range(player.Hand) {
		if card == move.Card {
			hasCard = true
		}
	}
	if !hasCard {
		err = errors.New(fmt.Sprintf("%s does not have card %s", move.Side, move.Card))
		return err, false
	}

	// are you following the trick, if able?
	if len(game.Trick) > 0 && move.Card.Suit != game.Trick[0].Suit {
		for _, card := range(player.Hand) {
			if card.Suit == game.Trick[0].Suit {
				err = errors.New(fmt.Sprintf("%s is able to follow trick (Trick=%s)", move.Side, game.Trick))
				return err, false
			}
		}
	}

	// if breaking hearts, are you legally able to?
	if !game.HeartsBroken && (move.Card.Suit == Hearts || move.Card == Card{Spades,Queen}) {
		for _, card := range(player.Hand) {
			if card.Suit != Hearts && (card != Card{Spades,Queen}) {
				err = errors.New(fmt.Sprintf("%s cannot break hearts yet", move.Side))
				return err, false
			}
		}
	}
	// all good
	return nil, true
}