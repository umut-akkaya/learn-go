package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)
	var s []int
	for number > 0 {
		digit := number % 10
		number = number / 10
		s = append(s, digit)
	}

	for _, i := range s {
		fmt.Print(i)
	}
	fmt.Println("")

}
