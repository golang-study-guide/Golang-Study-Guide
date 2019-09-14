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

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```

This file is our entry-point-file, for our executable program. It's best practice to name our entry-point-file, main.go. However the main.go file can get quite big, in which case you can break down the main.go into multiple smaller *.go files. In this scenario, the entry-point-file is the *.go file that contains the "func main()" function. By convention you can name this entry-point-file as:

- main.go
- you-golang-package-name.go, e.g. if your project creates a executable file called 'terraform', then this file is called terraform.go. In other word's it needs to match the repo name, e.g. if you're repo is https://github.com/Sher-Chowdhury/gsg_hello_world, then you should call it gsg_hello_world.go. 



. However all these new smaller *.go files needs to start with 'package main' as the [package declaration](https://golang.org/doc/code.html#PackageNames) so that golang still knows that they are part of the main package.   


Now you can run this code using the 'go run' command:

```bash
$ go run github.com/Sher-Chowdhury/gsg_hello_world
hello world
```

the [go run](https://golang.org/pkg/cmd/go/internal/run/) command specificially loads and runs all *.go files that have the first line set to  'package main'. Once all the 'package main' files are collated, `go run` then looks for the 'func main()' function which is the starting point of your code. 


Now if you want to run this module on another workstation, then you first need to commit and push this to github, and on another workstation, you just have to run:

```bash
$ go get github.com/Sher-Chowdhury/gsg_hello_world
$ go run github.com/Sher-Chowdhury/gsg_hello_world
hello world
```

That's why it's important to name your module after your github repo's url. 





