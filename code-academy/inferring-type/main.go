package main

import "fmt"

func main() {
  // Define daysOnVacation using := below:
  daysOnVacation := 7
  var hoursInDay = 24

  
  // Define hoursInDay using var and = below:
  
  
  fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")
}



/*
There is a way to declare a variable without explicitly stating 
its type using the short declaration := operator. 
We might use the := operator if we know what value we want our 
variable to store when creating it. For instance:
nuclearMeltdownOccurring := true
radiumInGroundWater := 4.521
daysSinceLastWorkplaceCatastrophe := 0
externalMessage := "Everything is normal. Keep calm and carry on."


Go also offers a separate syntax to declare a variable and infer its type. 
We could’ve written the same code from above as:

var nuclearMeltdownOccurring = true
var radiumInGroundWater = 4.521
var daysSinceLastWorkplaceCatastrophe = 0
var externalMessage = "Everything is normal. Keep calm and carry on."
*/
