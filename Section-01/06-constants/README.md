To understand how to write and use variables, take a look at the example code in the [gsg_constants](https://github.com/Sher-Chowdhury/gsg_constants)

There's 2 ways to run the code in this repo:


```bash
git clone https://github.com/Sher-Chowdhury/gsg_constants.git
cd gsg_constants
go run .
```

Or:

```bash
go get github.com/Sher-Chowdhury/gsg_constants
go run github.com/Sher-Chowdhury/gsg_constants
```

If you're running the `go get ...` command for the first time on you're workstation, then it ends up creating the `~/go` folder, and it then populates it:

```
$ cd ~/go
$ tree ~/go
/Users/sherchowdhury/go
├── bin
│   └── gsg_constants
└── src
    └── github.com
        └── Sher-Chowdhury
            └── gsg_constants
                ├── 01_initialised_only.go
                ├── 02_initialised_only_alternate_syntax.go
                ├── 03_initialised_only_alternate_syntax.go
                ├── 04_declared_variables\ copy.go
                ├── 05_initialise_then_declare_variable.go
                ├── 06_declared_variables_using_type_inference.go
                ├── 07_declared_variables_alternative_syntax.go
                ├── 08_declared_implied_variables.go
                ├── 09_declared_implied_variables_alt_syntax.go
                ├── 10_check_data_type.go
                ├── 11_the_pointer_datatype.go
                ├── go.mod
                └── main.go

5 directories, 14 files
```



Here we can see that the src code has been downloaded and a binary was automatically compiled from the source code. the `go run ...` then executes that binary. Alternatively you can run this binary directly by adding the bin folder to the PATH env variable:

```
$ PATH=$PATH:~/go/bin
$ gsg_constants
EG1 - Initialised only
'name' is set to:
'city' is set to:
'age' is set to: 0
...
```




