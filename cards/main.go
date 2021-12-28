package main

func main() {
	//var card string = "Ace of Spades"
	//above is long form and explicit, below does the same
	//go can infer types to a degree, only use := syntax when intially defining a var

	//declare a slice, so we can have a number of cards
	//cards := deck{"Ace of Diamonds", newCard()}
	//append does not modify original slice, but instead appends and creates a new slice
	//cards = append(cards, "Six of Spades")

	//i is index of card in the slice, here is a slice, take the slice and iterate over every single record in the slice
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

//func newCard() string {
//	return "Five of Diamonds"
//}

//Go is a statically typed language, so variables have an assigned type always
//Basic data types: bool, string, int, float64
//DS in Go:
//Array: Fixed length list of things, very fixed length of records, items inside must have identical type
//Slice: An array that can grow or shrink, items inside must have identical type
//Go is not OO, so no concept of classes
