package main

// we use the reflect module for finding out a variable's datatype. 
import (
  "fmt" 
  "reflect" 
)


// Golang can work out data type for you, so you don't need to explicitly specify the datatype
var (
  name = "Peter Parker" // string
  age = 18              // integer
  weight = 65.5         // float64
  is_a_student = true   // boolean
)

// here we find out what data type golang has determined for each variable
func main() {
  fmt.Println("Golang has decided to set the datatype for 'name' to:", reflect.TypeOf(name))
  fmt.Println("Golang has decided to set the datatype for 'age' to:", reflect.TypeOf(age))
  fmt.Println("Golang has decided to set the datatype for 'weight' to:", reflect.TypeOf(weight))
  fmt.Println("Golang has decided to set the datatype for 'is_a_student' to:", reflect.TypeOf(is_a_student))
}
