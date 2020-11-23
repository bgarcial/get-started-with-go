package main

import "fmt"

func main() {
	heistReady := true

	heistReady = false

	if heistReady {
		fmt.Print("Let's Go!")
	}

	isHungry := false
	if isHungry {
		fmt.Println("Eat the cookie")
	} else {
		fmt.Println("Step away from the cookie...")
	}

	// heistReady := false

	if heistReady {
		fmt.Println("Let's go!")
	} else {
		fmt.Println("Act normal.")
	}

}
