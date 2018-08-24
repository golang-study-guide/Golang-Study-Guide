package main

import (  
    "fmt"
)

func main() {  
    
    // an switch is aka a case statement. 

    name := "Peter Parker"
    switch name {
    case "Tony Stark":
        fmt.Println(name, "is Iron Man")
    case "Peter Parker":
        fmt.Println(name, "is Spiderman")
    case "Bruce Banner":
        fmt.Println(name, "is The Hulk")
    case "Steve Rogers":
        fmt.Println(name, "is Captain America")
    default:
        fmt.Println(name, "is not known")

    }
}
