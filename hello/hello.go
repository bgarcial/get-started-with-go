/*Declare a main package. 
In Go, code executed as an application must go 
in a main package.*/
package main

/*Importing 2 packages*/
import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// Configure the log package to print the command name ("greetings: ") 
	// at the start of its log messages, without a time stamp or 
	// source file information
	log.SetPrefix("greetings: ")
	log.SetFlags(0)


	// Get a greetings message and print it.
	// Calling function Hello on package greetings
	//Get a greeting by calling the greetings package’s 
	//Hello function.

	// Assign both of the Hello return values, including the error, to variables
	message, err := greetings.Hello("Bernardo")
	/*
	If an error was returned, print it to the console and
	exit the program
	*/
	if err != nil {
		log.Fatal(err)
		/*
		Use the functions in the standard library's log package 
		to output error information. If you get an error, you 
		use the log package's Fatal function to print the error and stop the program.
		*/
	}

	/*
	If no error was returned print the returned message to the
	console
	*/
	fmt.Println(message)
}


/*For this hello package is necessary to 
create a module 
> go mod init hello
go: creating new go.mod: module hello

go build command creates hello file
when we have arguments on greetings.Hello on hello.go.

If we do a change here, do again go build

*/