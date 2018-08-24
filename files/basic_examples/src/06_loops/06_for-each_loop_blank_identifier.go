package main

import (
  "fmt"
)


func main() {
  
  // Here we declare an array. 
  fruits :=  [3]string{"apple","banana","oranges"}


  // here we use '_' as a way to send the index value to /dev/null
  // this '_' is referred to as a blank identifier. 
  for _, array_value := range fruits {

    fmt.Println("the next value in the array is set to", array_value)

  }

 
}
