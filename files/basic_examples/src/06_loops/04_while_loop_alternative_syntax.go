package main

import (
  "fmt"
  "time"
)


func main() {


  for counter := 0; counter <= 10 ; counter++ {
    fmt.Println("counter is now: ", counter)
    time.Sleep(1 * time.Second)
  }

  // note when using this syntax, the counter variable only exists inside the for loop



/*
In this scenario, the for loop has 3 parts
for initialisation; condition; post {  
}


initialisation - something to run before starting the first loop run. 
condition - somehting that needs to evaluate to 'true' in order for next loop iteration to run
post - something to run as the final task of the current iteration. 

*/



}
