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

  fmt.Println("'name' is set to:", name, "and it's datatype is", reflect.TypeOf(name))
  fmt.Println("'city' is set to:", city, "and it's datatype is", reflect.TypeOf(city))
  fmt.Println("'age' is set to:", age, "and it's datatype is", reflect.TypeOf(age))
}
