package main

import (
  "fmt" 
  "reflect"
)



func main() {

  /*
  A slice is essentially uses an array behind the scenes, where we set the maximum number of values (cap) that the underlying array can hold,
  along with the size of the subset that will make up the original slice. 

  A slice is initialized using the make command. 
  */

  my_slice := make([]int, 5, 10)

  fmt.Println("datatype for my_slice is:", reflect.TypeOf(my_slice))
  fmt.Println("length for my_slice is:", len(my_slice))
  fmt.Println("cap  for my_slice is:", cap(my_slice))


  // if your slice, will have length==capped then you can write the shorthand

  my_slice2 := make([]int, 5)

  fmt.Println("datatype for my_slice is:", reflect.TypeOf(my_slice2))
  fmt.Println("length for my_slice is:", len(my_slice2))
  fmt.Println("cap  for my_slice is:", cap(my_slice2))

  // in fact an array is actually a declared slice where length==capped, therefore when creating a slice of this type, then can use exactly the same 
  // syntax as when creating an array, e.g. 

  my_slice3 := []int{10,20,30,40,50}
  fmt.Println("datatype for my_slice is:", reflect.TypeOf(my_slice3))
  fmt.Println("length for my_slice is:", len(my_slice3))
  fmt.Println("cap  for my_slice is:", cap(my_slice3))




}
