package main



import (
  "fmt"
  "errors"
)

// It's best practice for the last output to always be the error output. 
func divide(a int, b int) (c int, err error) {

  c = a / b 

  if b == 1 {  
    return 1, errors.New("dividing by 1 isn't allowed when using this function")
  }

  return c, nil
}

func main() {

  output_c1, err := divide(80, 10)
  if err != nil {
    fmt.Println("Encountered the following error message:", err)
  } else {
    fmt.Println(output_c1)
  }
  

  
  output_c2, err := divide(80, 1)
  if err != nil {
    fmt.Println("Encountered the following error message:", err)
  } else {
    fmt.Println(output_c2)
  }

}


