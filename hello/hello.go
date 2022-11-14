package main

import (
	"fmt"

	"rsc.io/quote"
)

// entry points
func main() {
	fmt.Println("Go Quote!")
	fmt.Println(quote.Go()) // Call external package
}
