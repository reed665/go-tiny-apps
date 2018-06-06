package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, cards := deal(cards, 5)

	fname := "my_deck"
	// err := hand.saveToFile(fname)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println("hand was saved to ", fname)

	fmt.Printf("Deck from file:\n")
	deck := newDeckFromFile(fname)
	deck.print()

	fmt.Printf("\nShuffled deck:\n")
	deck.shuffle()
	deck.print()
}
