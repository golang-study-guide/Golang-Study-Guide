package main

import (
  "fmt" 
  "reflect"
)

func main() {

  // here we didn't specify the datetype, so Golang works it out for us.
  var name = "Peter Parker" 
  var city = "New York"
  var age = 18

  fmt.Println("'name' is set to:", name)
  fmt.Println("'city' is set to:", city)
  fmt.Println("'age' is set to:", age)
}
