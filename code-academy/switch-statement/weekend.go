package main
import "fmt"

func main(){
	weekend := true
	dayOfWeek := 7
	//day := 2
 
	if !weekend {
	switch day {
	case 1: 
		fmt.Println("It's Monday already?")
	case 2: 
		fmt.Println("Tuesday... still better than Monday.")
	case 3:
		fmt.Println("Wednesday, it's hump day!")
	case 4:
		fmt.Println("Thursday, almost done.")
	case 5:  
		fmt.Println("Friday! TGIF!")
	default:
		fmt.Println("What day of the week is it?")
	}
	} else {
	fmt.Println("Take the day off!")
	}
}