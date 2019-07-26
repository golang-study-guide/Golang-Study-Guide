Installing Golang is really easy, just following the instructions in the <a href="https://golang.org/dl/">official guide</a>. For example, for RHEL/CentOSm we would do something like:


<pre>
golang_archive_filename=go1.10.3.linux-amd64.tar.gz
curl -o ${golang_archive_filename} https://dl.google.com/go/${golang_archive_filename}
tar -C /usr/local -xzf ${golang_archive_filename}
cd /usr/local/go/bin 
cp -rfp /usr/local/go/bin/* /usr/bin/
</pre>







However here are a few other ways to install Golang:


<h2>RHEL/CentOS install tips</h2>
A quick way to install Golang is via yum:

<pre>yum install golang</pre>

But note that there's a good chance the Go version yum installs could be at least several versions behind. 

<h2>MacOS install tips</h2>
You can install golang using <a href="https://brew.sh/">brew</a>. With brew it's as simple as:

<pre>
$ brew install go
</pre>

Then when new version of golang comes out, you can just do:

<pre>
$ brew upgrade golang
</pre>

<h2>Windows Install tips</h2>
You can install Golang using <a href="https://chocolatey.org/packages/golang">chocolatey</a>:

<pre>c:\ choco install golang</pre>

Then when new version of golang comes out, you can just do:

<pre>c:\ choco upgrade golang</pre>




<h2>Conclusion</h2>

As you can see, installing Go is quite straightforward. After it has been installed you can test if the install was successful by just running <code>go</code> on the command line. You an also check which version you have installed by running:

<pre>
$ go version
go version go1.10.3 darwin/amd64
</pre>