package main
 
import (
    "fmt"
)
 
func main() {



  for key, value := range map[string]int {
    "maths": 75,
    "english": 70,
    "chemistry": 68,
    "geography": 65,
  } {

    fmt.Println("The value for", key, "is", value)

  }

}

