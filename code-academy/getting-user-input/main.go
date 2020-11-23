package main

import "fmt"

/*
Another helpful method from the fmt package is .Scan() which allows us 
to get user input! 
*/

func main(){

	fmt.Println("How are you doing?") 

	var response1 string
	var response2 string 
	fmt.Scan(&response1)
	fmt.Scan(&response2)

	fmt.Printf("I'm %v %v", response1,response2)
	fmt.Println("\n***") 

	/*
	First, we print How are you doing? to the terminal. Then we declare a variable, 
	response with the type string. fmt.Scan(&response) takes the first value before 
	a space and stores it in response. In the terminal, we would see:

	How are you doing?

	If we type good

	How are you doing?
	good
	I'm good.
	*/

	/*
	However, if we tried to type in not bad:
	How are you doing?
	not bad
	I'm not.

	Only the not part is saved, since it was separated from bad by a space. 
	If we were expecting two values, we’d need to declare two variables ahead of time:
	
	fmt.Scan() expects addresses for arguments, hence the & before response1 and response2.
	
	*/

	fmt.Println("What would you like for lunch?")
  
  	// Add your code below:
  	var food string
  	fmt.Scan(&food)

  
  	fmt.Printf("Sure, we can have %v for lunch.", food)
}