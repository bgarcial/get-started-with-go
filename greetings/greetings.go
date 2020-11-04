package greetings
// Declare a greetings package to collect related functions.


import (
	"errors"
	"fmt"
)
// Hello returns a greeting for the named person and an error
// if string sent is empty or ""
func Hello(name string) (string, error) {

	// If no name was give, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
		// https://golang.org/pkg/errors/#example_New
	} 
	// If a name was received
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
	/*
	nil (meaning no error) as a second value in the successful return. 
	That way, the caller can see that the function succeeded. 
	We have to include two elements in return
	because Hello function take two parameters above
	go functions can return multiple values 
	https://golang.org/doc/effective_go.html#multiple-returns
	*/
}




/* https://golang.org/doc/tutorial/create-module
This function takes a name parameter whose type is string, 
and returns a string. 

----
In Go, a function whose name starts with a capital letter 
can be called by a function not in the same package. 
This is known in Go as an exported name.

Read about Exported names
https://tour.golang.org/basics/3
-----
Declare a message variable to hold your greeting.
In Go, the := operator is a shortcut for declaring and 
initializing a variable in one line (Go uses the value 
on the right to determine the variable's type). 
Taking the long way, you might have written this as:
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)

Use the fmt package's Sprintf function to create a greeting message. 
The first argument is a format string, and Sprintf substitutes the name 
parameter's value for the %v format verb. 
Inserting the value of the name parameter completes the greeting text.

Return the formatted greeting text to the caller.

We will call this function from another module
*/