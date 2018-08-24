package main

import (  
    "fmt"
)

func main() {


    num := 12 
    switch { 
    case num == 12:
        fmt.Printf("%d is also known as a dozen\n", num)
        // fallthrough lets you carry on after matched code is executed, until at least it encounter's another matching case block. 
        fallthrough
    case num < 5:
        fmt.Printf("%d is less than 100\n", num)
    case num < 100:
        fmt.Printf("%d is less than 100\n", num)
    case num < 200:
        fmt.Printf("%d is less than 200", num)
    }

}
