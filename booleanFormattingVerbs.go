// %t	Value of the boolean operator in true or false format (same as using %v)
package main

import (
	"fmt"
)

func main() {
	var i = true
	var j = false

	fmt.Printf("%t\n", i)
	fmt.Printf("%t\n", j)
}
