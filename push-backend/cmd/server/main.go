package main

import (
	"fmt"
)

func main() {
	deck := NewDeck()

	hand := []Card{
		deck[0],
		deck[1],
		deck[2],
		deck[3],
		deck[4],
		deck[5],
		deck[6],
		deck[7],
		deck[8],
		deck[9],
		deck[10],
	}

	fmt.Println("Hand:")
	for _, card := range hand {
		fmt.Printf("%s ", card.String())
	}
	fmt.Println()

	books := FindBooks(hand)
	fmt.Println("Books found in hand:")
	for i, book := range books {
		fmt.Printf("Book %d: ", i+1)
		for _, card := range book {
			fmt.Printf("%s ", card.String())
		}
		fmt.Println()
	}
	fmt.Println("Total books found:", len(books))

	runs := FindRunsWithWilds(hand)
	fmt.Println("Runs found in hand:")
	for i, run := range runs {
		fmt.Printf("Run %d: ", i+1)
		for _, card := range run {
			fmt.Printf("%s ", card.String())
		}
		fmt.Println()
	}
	fmt.Println("Total runs found:", len(runs))
}	
