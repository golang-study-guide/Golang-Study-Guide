package main

import (
  "fmt" 
  "reflect" 
)

// go comes with a bunch of builtin functions which are pre-imported
// https://golang.org/pkg/builtin/
var (
  price float64 = 237.00
)

func main() {
  fmt.Println("'price' is set to:", price)
  fmt.Println("'price' datatype is:", reflect.TypeOf(price))

  // here we use the 'int' built in function which converts a float to an integer. 
  rough_price := int(price)
  fmt.Println("'rough_price' is set to:", rough_price)
  fmt.Println("'rough_price' datatype is:", reflect.TypeOf(rough_price))
}
