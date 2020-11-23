package main
import "fmt"

func main() {
	switch pizzaTopping := "pepperoni" ; pizzaTopping {
	case "cheese": 
		fmt.Println("One pizza with extra cheese coming up!")
	case "pepperoni":
		fmt.Println("One pepperoni pizza in the oven!")
	default:
		fmt.Println("I don't have that topping, get something else.")
	}
}