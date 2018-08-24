package main

import (
  "fmt"
)

func main() {


  // here we added in an optional statement, 
  // which only exists in the scope of the if statement. 
  if small_number, average_number := 10, 50 ; small_number < average_number {
    fmt.Println(small_number, "is less than", average_number )
  }
   
  // the following statement won't work becuase it is outside the scope.   
  // fmt.Println(small_number, "is less than", average_number )


}
