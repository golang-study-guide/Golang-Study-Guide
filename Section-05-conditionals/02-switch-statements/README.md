```go
package main

type user struct {
  ID        int
  FirstName string
  Lastname  string
}


func identifyUser(aUser user) {

    switch aUser.FirstName {
    case "Homer":
        println("is the father")
    case "Marge":
        println("is the mother")
        fallthrough             // this executes the very next housewife block, irrespective if it matches.
    case "Housewife":
        println("works as a housewife")
    case "Bart":
        println("is the son")
    case "Lisa":
        println("is the daughter")
    default:
        println("Might be a recent addition to the family")
    }

}

func main() {
  user1 := user{
		ID:        01,
		FirstName: "Homer",
		Lastname:  "Simpson",
  }
  
  user2 := user{
		ID:        02,
		FirstName: "Marge",
		Lastname:  "Simpson",
  }

  user3 := user{
		ID:        03,
		FirstName: "Maggie",
		Lastname:  "Simpson",
  }

  identifyUser(user1)
  identifyUser(user2)
  identifyUser(user3)


}
```