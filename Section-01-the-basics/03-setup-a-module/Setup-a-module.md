Before you can start writing Golang code. You first need to set up a module. A module is just a folder that will store all your Golang project's source code. In my case I'll create a module called 'my_go_app1':


Create a new empty git repo and clone it. In my case I created and cloned:

```bash
$ git clone github.com:Sher-Chowdhury/gsg_hello_world.git
$ cd gsg_hello_world/
$ ls
```

At this stage this repo is empty. Next we run the "Go module initialisation" command inside this repo:


```bash
$ go mod init github.com/Sher-Chowdhury/gsg_hello_world.git
go: creating new go.mod: module github.com:Sher-Chowdhury/gsg_hello_world
```

Here I created a module, using the github repo's url as it's name. This ends up creating a file called "go.mod":

```bash
$ ls -l
total 8
-rw-------  1 schowdhury  staff  62  2 Sep 21:10 go.mod
$ cat go.mod 
module github.com/Sher-Chowdhury/gsg_hello_world

go 1.12
```

Then I created a file called "main.go":


```go
package main

func main() {
    println("hello world")
}
```

You can [run this here](https://play.golang.org/p/EGuGDgEM9Ex), or just run:

```bash
$ go run main.go
hello world
```

the [go run](https://golang.org/pkg/cmd/go/internal/run/) command specificially loads and runs all *.go files that have the first line set to  'package main'. Once all the 'package main' files are collated, `go run` then looks for the 'func main()' function which is the starting point of your code. 


You can create it as an executable binary and then run that:

```bash
$ go build main.go
$ ./main
hello world
```

Note: here the binary's name is taken from main.go, but without the .go extension. 


You can write all your go code into a single *.go file. However that single file can get quite big and difficult to read, which is why it's best practice to break your code into several smaller *.go files.

All the *.go files that forms part of your main go application, must have it's first line set as `package main`. This is the [package declaration](https://golang.org/doc/code.html#PackageNames) and this tells golang which *.go forms part of your main go application. Among these main files, exactly one file must contain the 'func main()' function. This file is that starting point of your go app. By convention this starting-point file is named as one of following:

- *main.go*
- *you-golang-package-name.go* - e.g. if your project creates a executable file called 'terraform', then this file is called terraform.go. In other word's it needs to match the repo name, e.g. if you're repo is https://github.com/Sher-Chowdhury/gsg_hello_world, then you should call it gsg_hello_world.go. 



A golang project is made up of [packages](https://golang.org/ref/spec#Packages), and each package in turn can make use of other packages. [println](https://golang.org/pkg/builtin/#println) is one of a [small handful of golang functions that comes builtin](https://golang.org/pkg/builtin) with golang. println is quite a primitive function, and that's why it's more common to use the more feature-rich alternative, Println, which comes as part of the fmt package:


```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```

Golang comes include with a [Standard Library](https://golang.org/pkg/#stdlib) of packages, of which the [fmt](https://golang.org/pkg/fmt/) package is one of them. You can get help info for fmt like this:

```bash
$ go doc fmt
package fmt // import "fmt"

Package fmt implements formatted I/O with functions analogous to C's printf
and scanf. The format 'verbs' are derived from C's but are simpler.
...

$ go doc fmt.Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

```



## downloading and running modules

Now if you want to run this module on another workstation, then you first need to commit and push this to github, and on another workstation, you just have to run:

```bash
$ go get github.com/Sher-Chowdhury/gsg_hello_world
$ go run github.com/Sher-Chowdhury/gsg_hello_world
hello world
```

You can also build the code from source:

```bash
$ go build .
$ ./gsg_hello_world
hello world
```


That's why it's important to name your module after your github repo's url. 





