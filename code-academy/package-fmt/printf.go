package main
import "fmt"

func main(){
	/*
	Using fmt.Println() and fmt.Print() we have the ability to concatenate strings, 
	i.e. combine different strings into a single string:
	*/
	guess := "C"
	fmt.Println("Is", guess, "your final answer?")
	// Prints: Is C your final answer?

	/*
	With fmt.Printf(), we can interpolate strings, or leave placeholders 
	in a string and use values to fill in the placeholders. 
	Let’s revisit the same example using fmt.Printf():
	*/
	//guess := "C"
	fmt.Printf("Is %v your final answer?", guess)
	// Prints: Is C your final answer?

	/*
	As long as we provide enough arguments, we can even add multiple placeholders:
	*/
	selection1 := "soup"
	selection2 := "salad"
	fmt.Printf("Do I want %v or %v?", selection2, selection1)
	// Prints: Do I want soup or salad?

	animal1 := "cat"
  	animal2 := "dog"
  
  	// Add your code below:
  	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

}