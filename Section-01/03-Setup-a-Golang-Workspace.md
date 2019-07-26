Before you can start writing Golang code. You first need to set up a Workspace. A workspace is just a folder that will store all your Golang related stuff. In my case I'll create a workspace called 'go_project':

<pre>
$ mkdir /root/go_project
</pre>

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
 