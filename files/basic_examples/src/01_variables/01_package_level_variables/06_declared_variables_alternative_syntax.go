package main

import (
  "fmt" 
)


// here multiple variables of the same data type can be declared on the same line
// although it makes it a little harder to read 
var (
  name, city = "Peter Parker", "New York"
  age = 18
)

func main() {
  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
