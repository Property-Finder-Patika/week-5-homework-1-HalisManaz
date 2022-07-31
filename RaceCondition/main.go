package main

import (
	"fmt"
)

var number = 2.0

func main() {
	fmt.Printf("Number at beginning %.2f\n", number)
	for transaction := 0; transaction <= 5; transaction++ {
		// Sequential programing
		dividing(transaction)
		multiply(transaction)

		// With concurrency
		// When program runs with concurrency after every end, it produces different result. That cause trouble

		/*
			To run program with concurrency remove the comment block
			go dividing(transaction)
			go multiply(transaction)
		*/
	}
	fmt.Printf("Number at the end %2.f\n", number)
}

// multiply function multiply given number with four
func multiply(transactionOrder int) {
	// Check before and after number
	fmt.Printf("%d: Number before multiplying: %.2f\n", transactionOrder, number)
	number = number * 4.0
	fmt.Printf("%d: Number after multiplying: %.2f\n", transactionOrder, number)
}

// dividing function dividing given number with two
func dividing(transactionOrder int) {
	// Check before and after number
	fmt.Printf("%d: Number before dividing: %.2f\n", transactionOrder, number)
	number = number / 2
	fmt.Printf("%d: Number after dividing: %.2f\n", transactionOrder, number)

}
