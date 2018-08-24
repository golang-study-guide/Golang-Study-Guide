package main

import (
  "fmt"
)


func main() {
  
  // Here we declare an array. 
  fruits :=  [3]string{"apple","banana","oranges"}


  // this 'range' syntax is used to iterate through a loop. the for+range statement is the equivalent to the for-each statement. 
  for array_key, array_value := range fruits {

    fmt.Println("the", array_key, "value in the array is set to", array_value)

  }

 
}
