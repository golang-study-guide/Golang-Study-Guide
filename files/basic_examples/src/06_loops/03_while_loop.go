package main

import (
  "fmt"
  "time"
)


func main() {

  counter := 0

  for counter <= 10 {
    
    fmt.Println("counter is now: ", counter)
    time.Sleep(1 * time.Second)
    counter = counter + 1
  }
}
