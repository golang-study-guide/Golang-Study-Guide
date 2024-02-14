Golang by default will executed lines in code in the order outlined in your code. 

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	fmt.Println("2")
    fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
    fmt.Println("7")
	fmt.Println("8")
}

```

E.g [this](https://go.dev/play/p/3ff0d7JQR0Z) will output:

```
1
2
3
4
5
6
7
8
```

That's because there's just one process (go routine) that runs by default. 

However you can have how multiple go routines. 


```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(10000 * time.Millisecond)
	fmt.Println(s)
}

func main() {
	go say("world")
	fmt.Println("hello")
}

```

In [this example](https://go.dev/play/p/dLBl1u5L60M) you might expect the output to be the following, after roughly a 10 second:

```
hello
world
```

However the output is the following after a 1 second wait:

```
hello
```

That's because the `main` function finished running and exited before `go say("world")` had a chance to complete.

If the main function contained a lot more code and it took about 2 mins to run, then we would have seen the output we hoped for originally, in which case, our go-routine optimised code would have run similar to how javascript runs code asynchronously by default.  


However you can make the main function wait for our goroutine to end, and that's done by using the `sync` packages's `WaitGroup`:


```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	time.Sleep(10000 * time.Millisecond)
	fmt.Println(s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say("world", &wg)
	fmt.Println("hello")
	wg.Wait()
}
```

In [this example](https://go.dev/play/p/aDSv-PIrgBu), the output is now:

```
hello
world
```

This time the world 'hello' appears straight away, and then it hangs for about 10 seconds before 'world' appears, and then it exits. 

In a way, waitgroup's are golang's version of javascript's async/await (promise.then) syntax.

Even if our `main` function did take several minutes, to run, it's still good to use waitgroups like this, just in case the main function completed faster than expected. 

