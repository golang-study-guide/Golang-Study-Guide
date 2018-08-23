package main

// we use the reflect module for finding out a variable's datatype. 
import (
  "fmt" 
  "reflect" 
)


// Golang can work out data type for you, so you don't need to explicitly specify the datatype
var (
  name = "Peter Parker"
  city = "New York"
  age = 18
)

// here we find out what data type golang has determined for each variable
func main() {
  fmt.Println("Golang has decided to set the datatype for 'name' to:", reflect.TypeOf(name))
  fmt.Println("Golang has decided to set the datatype for 'city' to:", reflect.TypeOf(city))
  fmt.Println("Golang has decided to set the datatype for 'age' to:", reflect.TypeOf(age))
}
