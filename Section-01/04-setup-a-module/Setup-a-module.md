Before you can start writing Golang code. You first need to set up a module. A module is just a folder that will store all your go project's source code. In my case I'll create a workspace called 'my_go_app1':


Create a new git repo and clone. In my case i created and cloned:

```bash
git clone github.com/Sher-Chowdhury/my_go_app1
```


run the Go module initialisation command inside this repo:


```bash
$ go mod init github.com/Sher-Chowdhury/my_go_app1
go: creating new go.mod: module github.com/Sher-Chowdhury/my_go_app1
```

Here I created a module, using the github repo's url as it's name. 

This ends up creating the file:

```bash
$ cat go.mod 
module github.com/Sher-Chowdhury/my_go_app1

go 1.12
```

Then I created a file called main.go:

```bash
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```

Now you can run this code a:

```bash
$ go run github.com/Sher-Chowdhury/my_go_app1
hello world
```

Now if you want to run this module in anther workstation, then i first need to commit and push this to github, and onother workstation, I just have to run:

```bash
$ go get github.com/Sher-Chowdhury/my_go_app1
$ go run github.com/Sher-Chowdhury/my_go_app1
hello world
```

That's why it's important to name your module after your github repo's url. 









A workspace needs to have 3 top level folders called <code>pkg</code>, <code>bin</code>, and <code>src</code>: 

<pre>
$ mkdir /root/go_project/{bin,pkg,src}
$ tree /root/go_project
/root/go_project
├── bin
├── pkg
└── src
</pre>

Next, you need to tell Go to use this workspace. That's done by setting an environment variable called 'GOPATH'. Let's first see what GOPATH is currently set to:

<pre>
$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
<strong>GOPATH="/root/go"</strong>
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build207657744=/tmp/go-build -gno-record-gcc-switches"

</pre>

I'm using bash, so I'll set this using the <a href="https://github.com/golang/go/wiki/SettingGOPATH#bash">.bash_profile</a> approach. 

You may want to have several workspaces for your various Go projects. If so, then you'll need to update your GOPATH each time you switch your focus to a different project. 


<pre>
$ export GOPATH=/my/current/project
</pre>


<h4>Good reads</h4>

https://golang.org/doc/code.html#Workspaces

https://github.com/golang/go/wiki/SettingGOPATH
 