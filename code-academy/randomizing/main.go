package main

import (
	"math/rand"
	"fmt"
)

func main() {
	fmt.Println(rand.Intn(100))
	/*
	we’re printing out a random number using rand and the Intn() method.
	it should print a random integer from 0 to 100. However, if we run our 
	program multiple times, we’ll find that it always prints 81. 
	*/

	// Edit amountLeft below: 
	amountLeft := rand.Intn(10000)
  
	fmt.Println("amountLeft is: ", amountLeft)
	
	  if amountLeft > 5000 {
		  fmt.Println("What should I spend this on?")
	} else {
	  fmt.Println("Where did all my money go?")
  }
}