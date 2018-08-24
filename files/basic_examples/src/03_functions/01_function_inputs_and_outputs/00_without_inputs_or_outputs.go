package main

import (
  "fmt"
  "time"
)


func sleep5() {


  time.Sleep(5 * time.Second)
  fmt.Println("5 second sleep is over")
}


func main() {

 sleep5()

}
