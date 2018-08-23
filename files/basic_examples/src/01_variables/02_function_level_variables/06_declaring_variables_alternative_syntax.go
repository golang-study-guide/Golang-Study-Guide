package main

import (
  "fmt" 
)

func main() {

  // Here we initialize only, ie. allocated storage space for storing the variable's value  
  var name, city string 
  var age int

  // Here we declare the value 
  name = "Peter Parker"
  city = "New York"
  age = 18

  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
