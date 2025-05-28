package main

import "fmt"

func main() {

	fmt.Println("This is a program to decide the number is even or not")

	fmt.Println("Enter the number:")
	var number int

	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println(err)
	} else {
		if number%2 == 0 {
			fmt.Println("Number is even")
		} else {
			fmt.Println("Number is odd")
		}
	}
}
