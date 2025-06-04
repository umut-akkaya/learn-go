package main

import (
	"fmt"
	"strconv"
)

func reverse(a int) int {

	var s []int
	var reversed_string string = ""
	for a > 0 {
		digit := a % 10
		a = a / 10
		s = append(s, digit)
		reversed_string = reversed_string + strconv.Itoa(digit)
	}

	reversed, err := strconv.Atoi(reversed_string)
	if err != nil {
		fmt.Println("Error")
	}
	return reversed
}

func main() {
	var number int
	fmt.Scanf("%d", &number)

	if number == reverse(number) {
		fmt.Println("This is a palindrome number")
	} else {
		fmt.Println("This is not a palindrome number")
	}
}
