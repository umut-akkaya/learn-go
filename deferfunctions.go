package main

import (
	"fmt"
)

func print_error() error {

	defer fmt.Println("Error catch 1")
	defer fmt.Println("Error catch 2")
	return fmt.Errorf("This is an error")
}

func main() {
	fmt.Println("Hello Friends")
	err := print_error()
	fmt.Println(err)
}
