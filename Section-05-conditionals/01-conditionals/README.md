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
		FirstName: "Marge",
		Lastname:  "Simpson",
  }

  if user1.Lastname == user2.Lastname {
    println("Users are related.")
  } else if user1.FirstName == "Homer" && user2.FirstName == "Marge" {
    println("There's a chance that both users are related through marriage")   
  } else {
    println("Users are not related.")
  }

}
```