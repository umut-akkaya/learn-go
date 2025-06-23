package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)

	fmt.Println(myArray[1:4])
	b := myArray[1:4]
	fmt.Println(b)

	for i := range b {
		b[i] = 0
	}
	fmt.Println(b)
	fmt.Println(myArray)
}
