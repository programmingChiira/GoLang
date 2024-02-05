// The Printf() function first formats its argument based on the given formatting verb and then prints them.
package main

import (
	"fmt"
)

func main() {
	var hello string = "Hello"
	var world string = "World"
	var number int = 12

	fmt.Printf("The first value is: %v, and the type is: %T\n", hello, hello)
	fmt.Printf("The second value is: %v, and the type is: %T\n", world, world)
	fmt.Printf("The second value is: %v, and the type is: %T\n", number, number)
}
