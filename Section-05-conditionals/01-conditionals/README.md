```go
package main

type user struct {
  ID        int
  FirstName string
  Lastname  string
}


func main() {
  user1 := user{
		ID:        01,
		FirstName: "Homer",
		Lastname:  "Simpson",
  }
  
  user2 := user{
		ID:        01,
		FirstName: "Homer",
		Lastname:  "Simpson",
  }

  if user1.Lastname == user2.Lastname {
    println("Users are related.")
  } else {
    println("Users are not related.")   
  }

}
```