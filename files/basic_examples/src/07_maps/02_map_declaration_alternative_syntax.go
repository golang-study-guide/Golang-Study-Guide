package main
 
import (
    "fmt"
)
 
func main() {



  my_exam_results := map[string]int {
    "maths": 75,
    "english": 70,
  }


  fmt.Println(my_exam_results)
  fmt.Println(my_exam_results["maths"])
  fmt.Println(my_exam_results["english"])
}

