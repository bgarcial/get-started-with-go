package main
import "fmt"

func main() {
	position := 2
 
	if position == 1 {
		fmt.Println("You won the gold!")
	} else if position == 2 {
		fmt.Println("You got the silver medal.")
	} else if position == 3 {
		fmt.Println("Great job on bronze.")
	} else {
		fmt.Println("Sorry, better luck next time?")
	}

	amountStolen := 64650

	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
  	} else if amountStolen >=5000 {
    	fmt.Println("Think of all the candy we can buy!")
  	} else {
		fmt.Println("Why did we even do this?")
	}
}