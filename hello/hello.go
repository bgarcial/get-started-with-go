/*Declare a main package. 
In Go, code executed as an application must go 
in a main package.*/
package main

/*Importing 2 packages*/
import (
	"fmt"
	"example.com/greetings"
)

func main() {
	// Get a greetings message and print it.
	// Calling function Hello on package greetings
	//Get a greeting by calling the greetings package’s 
	//Hello function.
	message := greetings.Hello("Bernardo")
	fmt.Println(message)
}


/*For this hello package is necessary to 
create a module 
> go mod init hello
go: creating new go.mod: module hello
*/