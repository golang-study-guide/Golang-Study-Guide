package main

import (
  "fmt"
)


func main() {


  for counter := 0; counter <= 10 ; counter++ {
    if counter ==  7 {
      continue
    }
    fmt.Println("counter is now: ", counter)
  }

}
