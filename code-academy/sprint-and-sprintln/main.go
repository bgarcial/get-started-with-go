package main

import "fmt"

func main(){
	/*
	We have other methods that don’t print strings, but format them 
	instead like fmt.Sprint() and fmt.Sprintln().
	*/
	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)
	// fmt.Sprint() doesn’t print out anything. Rather, it returned a value that we store in teacherSays.
 
	fmt.Print(teacherSays)
	// Prints: You scored a 100 on the test! Great job!
	fmt.Println("\n***") 
	/*
	fmt.Sprintln() works like fmt.Sprint() but it automatically 
	includes spaces between the arguments for us (just like fmt.Println() vs fmt.Print()!):
	*/
	quote := fmt.Sprintln("Look ma,", "no spaces!")
	fmt.Print(quote) // Prints Look ma, no spaces!
	/*
	Even though we didn’t add a trailing space in "Look ma," or 
	a leading space in "no spaces!", quote is concatenated with a space in between: 
	"Look ma, no spaces!".
	*/
	fmt.Println("\n***") 
	step1 := "Breathe in..."
  	step2 := "Breathe out..."
  
  	// Add your code below:
  	meditation := fmt.Sprintln(step1, step2)
	fmt.Print(meditation)
	// Prints Breathe in... Breathe out...
	  

}