The panic function is a builtin function that let's you exit the application if something happens that your
go app isn't designed:

```go
package main

func main() {
    println("attempt connection to DB")  // "println" comes builtin with golang, it is a more basic versio of fmt.Println()
    panic("failed to connect to db")
    println("perform CRUD operations with the DB")
}
```

this outputs:

```
attempt connection to DB
panic: failed to connect
```


also see: https://golang.org/doc/effective_go.html#panic



