package main

import ( 
	"io/ioutil"
	"math/rand"
	"fmt"
	"encoding/json"
	"log"
)

const filepath string = "./train/"
const trainprefix string = "mick"
const targetprefix string = "rock"
const batchsize int16 = 10
type FlatGame [53][3]bool
type FlatMove [52]bool

func CreateTrainingData(capacity int) {
	for capacity > 0 {
		mick, rock := new([batchsize]FlatGame)[:], new([batchsize]FlatMove)[:]
		for i := int16(0); i < batchsize; i++ {
			game := NewGame()
			for len(game.Trick) < 3 {
				hand := game.Players[game.ToPlay].Hand
				move := Move{
					game.ToPlay,
					hand[rand.Intn(len(hand))],
				}
				_, ok := game.ValidMove(move)
				if ok {
					game.Play(move)
				}
			}
			rst := game.Save()
			side := game.ToPlay
			bst := Card{ Clubs, Two }
			hiscore := game.Players[side].Total + 26
			hand := game.Players[side].Hand
			for _, card := range(hand) {
				m := Move{side, card}
				_, ok := game.ValidMove(m)
				if !ok {
					continue
				}
				game.Play(m)
				if game.Players[side].Total < hiscore {
					hiscore = game.Players[side].Total
					bst = card
				}
				game = rst
			}
			mick = append(mick, flattenGame(&game))
			rock = append(rock, flattenMove(Move{side, bst}))
		}
		writeToDisk(trainprefix, capacity, mick)
		writeToDisk(targetprefix, capacity, rock)
		capacity--
	}
}

func flattenGame (game *Game) (res FlatGame) {
	for _, c := range(game.Trick) {
		res[c.flatIdx()] = [3]bool{false,true,false}
	}
	for i, p := range(game.Players) {
		for _, c := range(p.Hand) {
			if Side(i) == game.ToPlay {
				if _, ok := game.ValidMove(Move{game.ToPlay, c}); ok {
					res[c.flatIdx()] = [3]bool{true,false,true}
					continue
				}
			}
			res[c.flatIdx()] = [3]bool{false,false,true}
		}
		for _, c := range(p.Points) {
			res[c.flatIdx()] = [3]bool{false,false,false}
		}
	}
	return
}

func flattenMove (move Move) (res FlatMove) {
	res[move.Card.flatIdx()] = true
	return
}

func ( card *Card ) flatIdx () int {
	return (int(card.Suit) * 13 + int(card.Value))
}

func writeToDisk (prefix string, id int, data interface{}) error {
	jdata, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return ioutil.WriteFile(fmt.Sprintf("%s%s%d", filepath, prefix, id), jdata, 0777)
}

func ( game *Game ) copyTotals () (res []int16) {
	for i, p := range(game.Players) {
		res[i] = p.Total
	}
	return
}

func ( game *Game ) Save () (g Game) {
	g = NewGame()
	for i, p := range(game.Players) {
		g.Players[i] = Player{
			p.Side,
			new([13]Card)[0:0],
			new([52]Card)[0:0],
			p.Total,
		}
		if len(p.Hand) > 0 {
			g.Players[i].Hand = append(g.Players[i].Hand, 
				p.Hand[0:len(p.Hand)-1]...)
		}
		if len(p.Points) > 0 {
			g.Players[i].Points = append(g.Players[i].Points, 
				p.Points[0:len(p.Points)-1]...)
		}
	}
	g.Trick = game.Trick[0:len(game.Trick)-1]
	return
}