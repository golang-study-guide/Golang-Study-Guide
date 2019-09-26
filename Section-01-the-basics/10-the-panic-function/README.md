The panic function is a builtin function that let's you exit the application if something happens that your
go app isn't designed:

```go
package main

func main() {

    var i int
	for {
        println("about to process: ", i) // "println" comes builtin with golang, it is a more basic versio of fmt.Println()
        if i == 3 {
            panic("I don't like number 3!")

        }

		i++
	}

}
```

this outputs:

```
attempt connection to DB
panic: failed to connect
```


also see: https://golang.org/doc/effective_go.html#panic



