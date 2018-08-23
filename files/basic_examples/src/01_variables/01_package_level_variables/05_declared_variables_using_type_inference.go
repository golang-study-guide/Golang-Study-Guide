package main

import (
  "fmt" 
)


// Golang can work out data type for you, so you don't need to explicitly specify the datatype
var (
  name = "Peter Parker"
  city = "New York"
  age = 18
)

func main() {
  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
