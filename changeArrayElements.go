package main

import (
	"fmt"
)

func main() {
	var elements = [3]int{53, 99, 121}

	elements[2] = 1000

	fmt.Println(elements)
}
