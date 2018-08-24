package main

import (
  "fmt"
  "time"
)


func main() {

  for counter := 0; counter <= 10 ; counter++ {
    if counter == 7 {
      fmt.Println("Lucky number 7 is reached")
      break
    }
    fmt.Println("counter is now: ", counter)
    time.Sleep(1 * time.Second)
  }

}
