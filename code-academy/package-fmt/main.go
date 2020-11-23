package main
import "fmt"

/*
fmt has a broader purpose, helping us format data, 
which is why it’s sometimes referred to as the 
format package (though it’s pronounced “fumpt”, and no that’s not a typo).
*/

func main(){
	/*
	In addition to Println(), there’s also Print() and Printf() 
	each with their own variation of how to print data. 
	There’s also Sprint(), Sprintln(), and Sprintf() which formats 
	but does not print anything to the terminal.
	*/
	// There's a mix of Println and Print!
	fmt.Println("Can", "you", "tell", "the", "difference?")
	fmt.Print("Between", "these", "two", "methods?")
	fmt.Print("Anything odd about", "the spacing? \n")
	fmt.Println("Don't worry if you can't spot it, we'll go through this together!")

	/*
	We can even get user input by using Scan()
	*/

	fmt.Print("The answer is", ": ")
	fmt.Print("12") 
	// The answer is: 12

	fmt.Println("Let's first see how", "the Println() method works.")
  	fmt.Println("Notice that each statement adds a newline for us.")
  	fmt.Println("There's also a default space", "between the string arguments.")
  
  	// Add your code below:
  	fmt.Print("Print ", "is ", "different")
	fmt.Print("See?")
	
	// Println() adds a line break for us.
	fmt.Println("Println", "formats", "really well")
	fmt.Println("Right?")  
}