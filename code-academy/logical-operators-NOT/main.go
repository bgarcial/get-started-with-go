package main
import "fmt"


func main() {
	/*
	not (!) operator. It negates (reverses) the value of a boolean.
	*/
	bored := true
	fmt.Println(!bored) // Prints false
 
	tired := false;
	fmt.Println(!tired) // Prints true

	readyToGo := true

	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}
