package main

import (
  "fmt" 
)

func main() {

  name, city, age := "Peter Parker", "New York", 18

  // here we omit the colon to mean that we are creating a new variable. Instead we are 
  // changing an existing variable
  age = 19

  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
