page 20
printing to standard output can be done using


fmt.Println()  // this adds newline at the end 
fmt.Print()    // this doesn't add newline

There's also:

fmt.Printf()

All three above are used to output to screen. 

This is like fmt.Print() but let's you do formatting, which in turn means write more code:

```
package main

import "fmt"

func main() {

	name := "Peter Parker"
	city := "New York"
	age := 18

        fmt.Print(name, city, age) // no spaces here
        fmt.Println(name, city, age)  // no spaces here, but new line
	fmt.Printf("my name is %s, I live in %s, and i am %d years old\n", name, city, age)

}
```

https://play.golang.org/p/1bK76lMVPhM





page 22

there's also Sprintf, Sprint, and Sprintln

These are used to construct string variable and store them in a another variable, e.g. 

```
package main

import "fmt"

func main() {

	name := "Peter Parker"
	city := "New York"
	age := 18

	sentence := fmt.Sprintf("my name is %s, I live in %s, and i am %d years old\n", name, city, age)

        fmt.Print(sentence)
}
```

https://play.golang.org/p/m_hsSuQCbKw


Similarly we have Fprintf, Fprint, and Fprintln



