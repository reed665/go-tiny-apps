package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
var cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func newDeck() deck {
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (hand deck, rest deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) String() string {
	s := []string(d)
	return strings.Join(s, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.String()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	str := string(bs)
	ss := strings.Split(str, ",")
	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i+1, c)
	}
}

func (d deck) contains(card string) bool {
	for _, c := range d {
		if c == card {
			return true
		}
	}
	return false
}
