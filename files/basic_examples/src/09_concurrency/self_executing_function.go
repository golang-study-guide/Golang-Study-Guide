package main

import (
  "fmt"
  "time"
  "sync"
)


func main() {

  // here we create a custom datatype called wait_for_end_of_goroutines
  var wait_for_end_of_goroutines sync.WaitGroup
  wait_for_end_of_goroutines.Add(2)

  go func() {
    defer wait_for_end_of_goroutines.Done()    // this runs just after it function gives a return value

    time.Sleep(5 * time.Second)
    fmt.Println("5 second sleep is over")
  }()

  go func() {
    defer wait_for_end_of_goroutines.Done()    // this runs just after it function gives a return value

    time.Sleep(1 * time.Second)
    fmt.Println("1 second sleep is over")
  }()

  wait_for_end_of_goroutines.Wait()   // this hangs until it get's to Done statements. 
}
