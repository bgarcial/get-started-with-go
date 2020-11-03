package main

import "fmt"

/* https://golang.org/doc/tutorial/getting-started#call  */
import "rsc.io/quote"

func main() {
    //fmt.Println("Hello, World")
    // https://pkg.go.dev/rsc.io/quote
    // https://github.com/rsc/quote/blob/v1.5.2/quote.go#L22
    fmt.Println(quote.Go())
}	
