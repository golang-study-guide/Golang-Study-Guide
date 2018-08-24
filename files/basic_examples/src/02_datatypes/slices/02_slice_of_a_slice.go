package main

import (
  "fmt" 
)



func main() {


  // here's a declared slice where I have set length==capped, aka an array. 
  my_slice := []int{10,20,30,40,50,60,70,80,90,100}

  slice_of_slice := my_slice[4:8]

  fmt.Println(slice_of_slice[1])  


}
