package main

func main() {
	//var card string = "Ace Of Spades"
	cards := newDeck()

	hand, remain := deal(cards, 4)
	hand.print()
	remain.print()

	//cards.print()

}
