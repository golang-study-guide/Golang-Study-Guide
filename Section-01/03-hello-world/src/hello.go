/*
A go code in a go project are usually organised into multiple *.go files.
to tell Go which *.go file is the starting point (i.e. main file) of your Go project, you need to add
to 'package main' line at the start of that file.
*/
package main

/*
Here we're importing the fmt package from the standard library
https://golang.org/pkg/
https://golang.org/pkg/fmt/
you can learn more about the fmt package by running:

$ go doc fmt
*/
import (
	"fmt"
)

/*
The 'main' function is used to tell Go, the first block of code to execute inside the main file
*/
func main() {
	/*
		You can learn more about Println by running:
		$ go doc fmt.Println
	*/
	fmt.Println("hello world")
}

/*
You can run this code using run:

go run hello.go


Behind the scenes, Go will compile this file into a binary and execute it. Then it deletes the binary
*/
