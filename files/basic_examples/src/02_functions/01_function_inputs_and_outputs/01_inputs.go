package main

import (
  "fmt"
)


// here's how to specify two input arguments
func plus(a int, b int) {

  c := a + b
  fmt.Println("The 'a' parameter is set to:", a)
  fmt.Println("The 'b' parameter is set to:", b)
  fmt.Println("The total is:", c)
}


func main() {

  plus(1, 2)

  // Note, the following line won't work, because a variables scope is limited to it's own function
  // fmt.Println("The total is:", c)
}
