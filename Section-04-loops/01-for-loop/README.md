In golang, there is only the for-loop. 

I.e. there are no such thing as do-loops, while-loops, until-loops,..etc. 


There are 4 types of for-loops, 

- loop till condition
- loop till condition with post clause
- infinite loop
- loop over collections

## infinite loop

```go
package main
import (
    "fmt"
)

func main() {
    var i int  // this defaults to 0
    for i < 5 {
        fmt.Println("i is set to: ", i)      
    }
}
```

This outputs:

```
i is set to:  0
i is set to:  0
i is set to:  0
.
.
...etc
```


## loop till condition

```go
package main

func main() {
    var i int
    for i < 5 {
        println(i) 
        i++      
    }
}
```

This outputs:

```
0
1
2
3
4
```

## exit loop early using "break"


```go
package main

func main() {
    var i int
    for i < 5 {
        println(i)
        if i == 2 {
          break
        }
        i++      
    }
}
```

This outputs:

```
0
1
2
```

## skip loop early using "continue"

```go
package main

func main() {
    var i int
    for i < 5 {
        i++
        if i == 2 {
          continue
        }
        println(i)
    }
}
```

this outputs:

```
0
1
3
4
5
```


## loop with clauses 

[loop with clauses](https://tour.golang.org/flowcontrol/1)

this loop has 3 parts to it's definition:

```go
package main

func main() {

	for i := 0; i < 5; i++ {
		println(i)
	}
	
}
```

In this scenario, the variable i only exists inside the scope of the for loop. 

this outputs:

```
0
1
3
4
```

## Infinite loop (ugly syntax)

This leaves the middle 'conditional' segment, (and optionally the other 2 statements) blank:

```go
package main

func main() {
  var i int
	for ; ; {
		println(i)
		i++
	}
	
}
```

this outputs:

```
0
1
2
3
4
5
6
7
...etc
```

## Infinite loop (pretty syntax)

Here we remove the "; ;" syntax and initialize the counter variable outside the loop:

```go
package main

func main() {
    var i int
	for {
		println(i)
		i++
	}
	
}
```

this outputs:

```
0
1
2
3
4
5
6
7
...etc
```

## Looping over collections

looping over slices:

```go
package main

func main() {
        mySlice := []string{"a", "b", "c"}
	for i:=0; i<len(mySlice); i++ {
		println(mySlice[i])
	}
	
}
```

this outputs:

```
a
b
c
```

Here's a more neater syntax using 'range':

```go
package main

func main() {
        mySlice := []string{"a", "b", "c"}
	for myIndex, myValue := range(mySlice) {
		println(myIndex, myValue)
	}
	
}
```

This outputs:

```
0 a
1 b
2 c
```

iterating over a map:

```go
package main

func main() {
        pets := map[string]int{"dog": 4, "cat": 7, "hamster": 1}
	for typeOfPet, ageOfPet := range(pets) {
		println(typeOfPet, ageOfPet)
	}
	
}
```

this outputs:

```
dog 4
cat 7
hamster 1
```

If you're only interested in the map's value, you can just do:

```
package main

func main() {
        pets := map[string]int{"dog": 4, "cat": 7, "hamster": 1}
	for _, ageOfPet := range(pets) {
		println(ageOfPet)
	}
	
}
```

This outputs:

```
4
7
1
```

If however you're only interested in the map's key, you can just do:

```
package main

func main() {
        pets := map[string]int{"dog": 4, "cat": 7, "hamster": 1}
	for typeOfPet, _ := range(pets) {
		println(typeOfPet)
	}
	
}
```

or using an even more shorthand syntax, you can do:

```go
package main

func main() {
        pets := map[string]int{"dog": 4, "cat": 7, "hamster": 1}
	for typeOfPet := range(pets) {
		println(typeOfPet)
	}
	
}
```


both of these examples, outputs:

```
dog
cat
hamster
```