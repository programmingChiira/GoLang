package main

import (
	"fmt"
)

func myFunction(name string, age int) {
	fmt.Println("Hi there, my name is ", name, "I am ", age)
}

func main() {
	myFunction("Dennis", 24)
}
