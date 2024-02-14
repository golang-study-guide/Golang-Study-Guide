This is a really good course - https://www.pluralsight.com/courses/go-programming-concurrent

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




## Channels

Channels offers a way to send data (via an input parameter) to a go routine to process:



```go
package main

import (
	"fmt"
	"time"
)

type product struct {
	name     string
	quantity int
	price    float32
}

func processShoppingList(in chan product) {

	for {
		item, ok := <-in
		if !ok {
			fmt.Println("Channel closed!")
			break
		}
		fmt.Printf("Need to buy %d %s for £%.2f each.\n", item.quantity, item.name, item.price)
	}

}

func main() {

	shoppingList := make(chan product)

	go processShoppingList(shoppingList)

	item1 := product{
		name:     "skittles",
		quantity: 4,
		price:    2.50,
	}
	shoppingList <- item1

	time.Sleep(3000 * time.Millisecond)
	item2 := product{
		name:     "bananas",
		quantity: 2,
		price:    1.50,
	}
	shoppingList <- item2

	time.Sleep(3000 * time.Millisecond)
	close(shoppingList)

	time.Sleep(3000 * time.Millisecond)
}

```

[This example](https://go.dev/play/p/O7HrtZFNCFH) gives the output:

```
$ go run main.go
Need to buy 4 skittles for £2.50 each.   # a few seconds pause.   
Need to buy 2 bananas for £1.50 each.    # a few seconds pause. 
Channel closed!                          # a few seconds pause.
```

Some key points:

- We've picked the `processShoppingList` function to run as a goroutine. This function's input parameter is a variable of the type channel (that accepts `product` custom types onto it's queue).
-  This `processShoppingList` goroutine, is listening for one `product` object to get added onto the channel (queue), and as soon as that happen's the function gets executed, and then the goroutine exits. It exits because, the channel has been emptied out (drained). That means that Go throws an `fatal error: all goroutines are asleep - deadlock!` error when item2 gets added to the channel. To prevent an early exit, we added a for-loop, so that the function remains running as long as the channel is open. 
- We ended up running the `processShoppingList` twice, once for item1, and once for item2, without expeclitly calling the function. We just simply added `product` objects to the channel. 



Here's [another example](https://gobyexample.com/channels):

```go
package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { 
		messages <- "ping" 
	}()

    msg := <-messages
    fmt.Println(msg)
}

```

This time, The goroutine is an iife (immediately invoked function expression) that puts an item onto the channel, rather than taking anything of the channel. Then the `main` itself takes things off the channel and processes them. Here's what the output looks like:

```bash
$ go run main.go
ping
```

Channels are specially designed to let goroutines communicate with each other. 

Also goroutines uses very little RAM, so you should use them a lot. 


## Buffered channels

Let's take a look at [this example](https://go.dev/play/p/KMKHc4VLFqC):

```go
package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func sender() {
	ch <- "message"
	fmt.Println("item added to the channel")
}

func receive() {
	time.Sleep(3000 * time.Millisecond)
	message := <-ch
	fmt.Println("Received this messsage from the channel: " + message)
}

func main() {
	go sender()
	go receive()
	time.Sleep(10000 * time.Millisecond)

}
```

After 3 seconds, this outputs:

```
Received this messsage from the channel: message
item added to the channel
```

The sender function get's blocked and doesn't run it's fmt line, until the receive function takes the item off the channel. This is handy if you want to keep the 2 goroutines in sync. 

However, you can set a cache size (buffer) in your channel, so that it sender can continue working by filling up the available, cache. Once the cache is full (which happens if the takes items off the channel to slowly), then the blocking starts happening again. To set a cache, you just defined the cache size when declare+initialising the channel variable, e.g:

```go
var ch = make(chan string, 1)
```

Here's [another example](https://go.dev/play/p/Mb_BFMLWOfy) (this time not using any goroutines):

```go
package main

import (
	"fmt"
)

func main() {
	var ch = make(chan string)

	ch <- "message"

	fmt.Println(<- ch)	
}
```

this outputs the error:

```
fatal error: all goroutines are asleep - deadlock!
```

That's because we never reach the print line, because the `main` function is blocked as soon an item is put onto the channel. 

## Directional channels


```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)

	go func(ch chan string) {
		fmt.Println(<-ch)
	}(ch)

	go func(ch chan string) {
		ch <- "message"
	}(ch)

	time.Sleep(1000 * time.Millisecond)

}
```

This outputs:

```
go run main.go
message
```