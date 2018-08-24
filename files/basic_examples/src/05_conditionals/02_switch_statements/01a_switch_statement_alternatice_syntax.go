package main

import (  
    "fmt"
)

func main() {  
    
    // an switch is aka a case statement. 

    name := "Peter Parker"
    switch {
    case name == "Tony Stark":
        fmt.Println(name, "is Iron Man")
    case name == "Peter Parker":
        fmt.Println(name, "is Spiderman")
    case name =="Bruce Banner":
        fmt.Println(name, "is The Hulk")
    case name == "Steve Rogers":
        fmt.Println(name, "is Captain America")
    default:
        fmt.Println(name, "is not known")

    }
}
