package main

import "fmt"

func main(){

	/*
	when we need to use the original value of a variable for a calculation 
	(or any general manipulation) and then update the variable to store the 
	newly calculated value. Let’s say we were keeping track the cost of items 
	in our grocery basket:
	*/
	var basketTotal float64
	carrotPrice := 0.75

	basketTotal = basketTotal + carrotPrice
	fmt.Println(basketTotal) // 0.75

	/*
	Notice that we used the original value of basketTotal which wasn’t assigned a value, 
	so it has a default value of 0.0, added carrotPrice (0.75) and then assigned the 
	computed value to basketTotal. If we add another item:
	*/
	spinachPrice := 1.50
	basketTotal = spinachPrice + basketTotal // 2.25
	/*
	This time, we added spinachPrice to basketTotal and stored the new value again in 
	basketTotal, thereby updating our running total!
	*/

	/*
	Updating a variable by adding another number to itself 
	and saving the new value is so common that Go has a 
	shorthand for it, the += operator. We could have done 
	the same operation using the following syntax:
	*/
	spinachPrice := 1.50
	basketTotal += spinachPrice // Print 2.25
	/*
	Notice that basketTotal = spinachPrice + basketTotal 
	and basketTotal += spinachPrice do the same thing! 
	We can also do the same for strings (i.e. concatenating strings together):
	*/
	command := "Hold my "
	beverage := "soda"

	command += beverage
	fmt.Println(command) // Prints: Hold my soda


}