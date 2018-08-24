package main
 
import (
    "fmt"
)
 
func main() {


  // maps (aka dictionary, aka hash table) are initialised and declared using the 'make' builtin function
  // maps are 'reference types', i.e. when you feed a map into a function, then it doesn't get cloned. instead maps original content will get modified.  



  my_exam_results := make(map[string]int)

  // this says that the map's keys will be in the form of strings, whereas the value will be in the form of integers. 

  my_exam_results["maths"] = 75
  my_exam_results["english"] = 70


  fmt.Println(my_exam_results)  
  fmt.Println(my_exam_results["maths"])
  fmt.Println(my_exam_results["english"])


}

