package main

import (
  "fmt" 
)



func main() {


  // here's a declared slice where I have set length==capped, aka an array. 
  my_slice := []int{10,20,30,40,50,60,70,80,90,100}

  slice_of_slice := my_slice[6:9]

  slice_of_slice = append(slice_of_slice, 200)
  fmt.Println(slice_of_slice)  
  fmt.Println(slice_of_slice)  
  fmt.Println(len(my_slice))  

  slice_of_slice = append(slice_of_slice, 300)
  fmt.Println(slice_of_slice)  
  fmt.Println(slice_of_slice)  
  fmt.Println(len(my_slice))  

  slice_of_slice = append(slice_of_slice, 400)
  fmt.Println(slice_of_slice)  
  fmt.Println(slice_of_slice)  
  fmt.Println(len(my_slice))  


  slice_of_slice = append(slice_of_slice, 500)
  fmt.Println(slice_of_slice)  
  fmt.Println(slice_of_slice)  
  fmt.Println(len(my_slice))  


  // if you reach the cap limit then go will just double the cap and keep doubling it behind the scenes. 
  // it does this by cloneing existing array. 

}
