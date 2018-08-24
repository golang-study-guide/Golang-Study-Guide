package main

import (  
    "fmt"
)

func main() {  
    

    // the name variable is only available inside the switch function. 
    switch name := "Peter Parker"; name {
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
