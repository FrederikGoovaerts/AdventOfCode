package main

import (
	"fmt"
)

const modVal = 20201227

func getLoopNb(public int) int {
	curr := 1
	loop := 0
	for curr != public {
		loop++
		curr = (curr * 7) % modVal
	}
	return loop
}

func doLoop(sub int, loop int) int {
	curr := 1
	for i := 0; i < loop; i++ {
		curr = (curr * sub) % modVal
	}
	return curr
}

func main() {
	// ex1
	// cardPub := 5764801
	// doorPub := 17807724

	// input
	cardPub := 8184785
	doorPub := 5293040

	cardLoop := getLoopNb(cardPub)

	fmt.Println(doLoop(doorPub, cardLoop))
}
