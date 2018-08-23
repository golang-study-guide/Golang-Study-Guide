package main

import (
  "fmt" 
  "os" 
)

func main() {
 

  // the following prints out values that gets listed when running 'env' in bash terminal
  fmt.Println(os.Getenv("HOME"))
  fmt.Println(os.Getenv("HOSTNAME"))
}
