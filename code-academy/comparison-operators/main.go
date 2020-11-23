package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

	if lockCombo == robAttempt {
		fmt.Println("The vault is now opened.")
	}
	vaultAmt := 2356468
	if vaultAmt >= 200000 {
		fmt.Println("We're going to need more bags")
	}
	rightTime := true
	rightPlace := true

	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	enoughRobbers := false
	enoughBags := true

	// Checking only one option
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
	  
}