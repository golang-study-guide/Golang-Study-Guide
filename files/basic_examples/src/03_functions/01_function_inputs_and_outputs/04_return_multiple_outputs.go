package main

import (
  "fmt"
)

// returning multiple outputs is done in a similar way to inputs
func plus(a int, b int) (c int, d int) {
  c = a + b 
  d = a * b
  return c, d
}

func main() {
  fmt.Println(plus(17, 10))


  // since the function has 2 outputs we capture them in seperate variables
  output_c, output_d := plus(8, 10)

  fmt.Println(output_c)
  fmt.Println(output_d)
}


