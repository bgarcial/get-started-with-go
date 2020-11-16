package main

import "fmt"

/*
Go actually allows us to declare multiple variables on a single line, 
in fact, there’s a few different syntaxes!
*/

func main(){

	var part1, part2 string
	part1 = "To be..."
	part2 = "Not to be..."
	fmt.Println(part1,part2)

	/*
	If we already know what values we want to assign our variables we can use := like so:
	*/
	quote, fact := "Bears, Beets, Battlestar Galactica", true
	fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica
	fmt.Println(fact) // Prints: true

	var magicNum, powerLevel int32
  	magicNum = 2048
  	powerLevel = 9001
  
  	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
  
  	// Define amount and unit below:
  	amount, unit := 10, "doll hairs"
  
  	fmt.Println(amount, unit, ", that's expensive...")
}