package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &number)
	var sum int
	sum = 0

	for number > 0 {
		digit := number % 10
		sum = digit + sum
		number = number / 10
	}

	fmt.Println("Sum of the digits is ", sum)

}
