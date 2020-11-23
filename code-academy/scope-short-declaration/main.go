package main
import "fmt"

func main() {
	x := 8
	y := 9
	/*
	we first declare product. Notice that product is separated from the condition using a semi-colon ;
	*/
	if product := x * y; product > 60 {
  		fmt.Println(product, "  is greater than 60")
	}

	
	/*
	when using the short variable declaration in if or switch statements is that the 
	declared variable is scoped to the statement blocks.
	Declaring season within switch statement
	*/
	switch season := "summer" ; season {
	case "summer":
	  fmt.Println("Go out and enjoy the sun!")
	}

	//success := true

	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000
	//numOfThieves := 5

	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
	  fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
	  fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}



}
