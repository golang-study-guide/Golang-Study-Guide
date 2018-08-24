package main

import (
  "fmt" 
  "reflect"
)



func main() {


    // here we defined an array that can hold 3 items, all of which have to be of the string datatype. 
    // note, you can't change the number of items, afterwards, e.g. you can't convert fruits to [4]string
    // Because of this restriction, arrays are rarely used in Golang. 
    // The alternative is to use slices. Slices are similar to arrays but are more flexible. 
    var fruits [3]string
    

    fruits[0] = "apple"
    fruits[1] = "banana"
    fruits[2] = "oranges"

    fmt.Println(fruits)
    fmt.Println(reflect.TypeOf(fruits))
    fmt.Println(fruits[1])
    fmt.Println("length of fruits is",len(fruits))   // len is a builtin golang function

}
