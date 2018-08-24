package main

import (
  "fmt"
)

func main() {

  // The condition must evaluate to a boolean datatype value. 

  small_number := 10
  average_number := 50


  if small_number == average_number {
    fmt.Println(small_number, "is equal to", average_number )
  } else if average_number > small_number {
    fmt.Println(small_number, "is is less than", average_number )
  } else {
    fmt.Println(small_number, "is greater than", average_number )
  }

   
}
