package main

import (
  "fmt"
  "time"
)


func main() {


  // note this function doesn't have a name. 
  // also round brackets needs to be placed just after closing curly brace, in order to make it self executing. 
  func() {
    time.Sleep(5 * time.Second)
    fmt.Println("5 second sleep is over")
  }()

}
