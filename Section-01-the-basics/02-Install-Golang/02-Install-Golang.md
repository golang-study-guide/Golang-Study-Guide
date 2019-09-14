Installing Golang is really easy, just following the instructions in the <a href="https://golang.org/dl/">official guide</a>. For example, for RHEL/CentOSm we would do something like:


```
golang_archive_filename=go1.10.3.linux-amd64.tar.gz
curl -o ${golang_archive_filename} https://dl.google.com/go/${golang_archive_filename}
tar -C /usr/local -xzf ${golang_archive_filename}
cd /usr/local/go/bin 
cp -rfp /usr/local/go/bin/* /usr/bin/
```






However here are a few other ways to install Golang:


## RHEL/CentOS install tips
A quick way to install Golang is via yum:

```
yum install golang
```

But note that there's a good chance the Go version yum installs could be at least several versions behind. 

## MacOS install tips
You can install golang using <a href="https://brew.sh/">brew</a>. With brew it's as simple as:

```
$ brew install go
```

Then when new version of golang comes out, you can just do:

```
$ brew upgrade golang
```

## Windows Install tips
You can install Golang using [chocolatey](https://chocolatey.org/packages/golang)

```c:\ choco install golang```

Then when new version of golang comes out, you can just do:

```c:\ choco upgrade golang```







## Conclusion

As you can see, installing Go is quite straightforward. After it has been installed you can test if the install was successful by just running `go` on the command line. You an also check which version you have installed by running:

```
$ go version
go version go1.10.3 darwin/amd64

$ go help
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         download and install packages and dependencies
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildmode   build modes
	c           calling between Go and C
	cache       build and test caching
	environment environment variables
	filetype    file types
	go.mod      the go.mod file
	gopath      GOPATH environment variable
	gopath-get  legacy GOPATH go get
	goproxy     module proxy protocol
	importpath  import path syntax
	modules     modules, module versions, and more
	module-get  module-aware go get
	packages    package lists and patterns
	testflag    testing flags
	testfunc    testing functions

Use "go help <topic>" for more information about that topic.
```