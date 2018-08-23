package main

import (
  "fmt" 
)

func main() {

  // here's a more shorthand syntax, where multiple variables of the same syntax are initialised on the same line
  // althought it makes it a litter harder to read
  var name, city string 
  var age int

  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
