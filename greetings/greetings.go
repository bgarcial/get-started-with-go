package greetings

// Declare a greetings package to collect related functions.

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person and an error
// if string sent is empty or ""
func Hello(name string) (string, error) {

	// If no name was give, return an error with a message
	if name == "" {
		return name, errors.New("empty name")
		// https://golang.org/pkg/errors/#example_New
	}
	// If a name was received
	// Return a greeting that embeds the name in a message
	// and the message is created using a random format
	message := fmt.Sprintf(randomFormat(), name)
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

// init sets values for variables used in the function
// Add an init function to seed the rand package with the current time
// Go executes init functions automatically at program startup, 
//after global variables have been initialized.
/*
In Hello, call the randomFormat function to get a format for the 
message you'll return, then use the format and name value 
together to create the message.
Return the message (or an error) as you did before.
Your hello.go needn't change.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
Add a randomFormat function that returns a randomly selected 
format for a greeting message. 
Note that randomFormat starts with a lowercase letter, 
making it accessible only to code in its own package 
(in other words, it's not exported).
*/

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	/* A slice of messages format
	Declare a formats slice with three message formats. 
	When declaring a slice, you omit its size in the brackets, 
	like this: []string. This tells Go that the array underlying 
	a slice can be dynamically sized.
	*/
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v! ",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	// generate a random number for selecting an item from the slice.
	return formats[rand.Intn(len(formats))]
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
