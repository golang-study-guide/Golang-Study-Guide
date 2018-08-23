package main

import (
  "fmt" 
)


// you can declare multiple variables of the same type on a single line. 
// although this can make the code harder to read 
var (
  name, city string
  age  int
)

func main() {
  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
