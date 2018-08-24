package main

import (
  "fmt"
)


func main() {


  for counter := 0; counter <= 10 ; counter++ {
    if counter ==  7 {
      fmt.Println("Lucky number 7 is reached")
      break
    }
    fmt.Println("counter is now: ", counter)
  }

}


// note: you can break out of mulitiple nested loops by labelling your breaks. 
