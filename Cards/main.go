package main

import "fmt"

var num int // declare outside of func
var count int = 3

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" // compiler figures out the data type even though go is static typed
	fmt.Println(card)
	card = "King" // colon only needed for init
	fmt.Println(card)

	//var num int
	num = 10
	fmt.Println(num)
	fmt.Println(count)
	// card from func
	card2 := newCard()
	fmt.Println(card2)

	printIt("Hello schmello")
}

// return types
func newCard() string {
	return "Queen"
}
