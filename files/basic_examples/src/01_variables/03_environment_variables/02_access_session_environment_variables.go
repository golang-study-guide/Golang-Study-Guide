package main

import (
  "fmt" 
  "os" 
)

func main() {
 

  // The following only works if you feed in value as part of a run, e.g.
  // $ FOO=BAR go run 02_access_session_environment_variables.go 
  fmt.Println(os.Getenv("FOO"))
}
