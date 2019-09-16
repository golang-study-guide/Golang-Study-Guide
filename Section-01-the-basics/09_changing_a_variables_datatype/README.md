You can change a variable's datatype by using the T(v) syntax, where "T" is string, float64, ...etc, and 'v' is the variable. e.g.:



```
package main

import (
	"fmt"
    "reflect"
)

func main() {
	var x int = 2
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(x)
	xString := float64(x)
	fmt.Println(reflect.TypeOf(xString))
	fmt.Println(xString)
}
```