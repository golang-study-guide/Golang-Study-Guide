package main

import (
  "fmt"
)


// since they are both the same datatype, you can shorten it to just: 
func plus(a, b int) {

    c := a + b
    fmt.Println("The 'a' parameter is set to:", a)
    fmt.Println("The 'b' parameter is set to:", b)
    fmt.Println("The total is:", c)
}


func main() {

    plus(1, 2)
}
