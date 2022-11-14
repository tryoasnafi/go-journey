package main

import "fmt"

func main() {
    // var name type = expression
    var firstName string = "Tryo"

    // shorthand (type inference)
    lastName := "Asnafi"

    fmt.Println("Halo,", firstName, lastName)
    
    // multiple declarations
    first, second, third := "First", 2, true

    fmt.Println(first)
    fmt.Println(second)
    fmt.Println(third)

    // blank identifier (_)
    quote, _ := "Golang is", "concise"
    fmt.Println(quote)
    // fmt.Println(_) >>>>> cannot use _ as value or type
    
    // keyword new for reference pointer
    point := new(int8)
    fmt.Println(point) // output: memory address
    fmt.Println(*point) // dereference using asterisk (*)

}
