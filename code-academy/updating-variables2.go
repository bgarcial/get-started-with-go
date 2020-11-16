package main

import "fmt"

/*
In addition to += (yes, pun intended), Go has other arithmetic operations 
that perform calculations and update the variable’s value:

-= to subtract from the variable.
*= to multiply the variable by a factor.
/= to divide the variable by a dividend.
Let’s get some practice using these shorthand operators.
*/

func main() {
  coolSneakers := 65.99
  niceNecklace := 45.50
  
  // Define taxCalculation here
  var taxCalculation float64
  // Add coolSneakers to taxCalculation
  taxCalculation += coolSneakers
  fmt.Println("The sneakers value was added to calculate his taxes later on ", taxCalculation)

  // Add niceNecklace to taxCalculation
  taxCalculation += niceNecklace 
  fmt.Println("The necklace value was added to calculate his taxes later on ", taxCalculation)


  // Compute the NYC sales tax
  // 8.875% of the purchase here:
  taxCalculation *= .08875
  // Uncomment this line for a receipt!
  fmt.Println("Purchase of", coolSneakers + niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers + niceNecklace + taxCalculation)
}