package main
 
import (
    "fmt"
)
 
func main() {



  my_exam_results := map[string]int {
    "maths": 75,
    "english": 70,
    "chemistry": 68,
    "geography": 65,
  }


  for key, value := range my_exam_results {

    fmt.Println("The value for", key, "is", value)

  }
  // note the output ordering can be different on successive runs because maps stores key-value pairs in random sequence. 

}

