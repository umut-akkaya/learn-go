package main

import "fmt"

func main() {

	a := [3]int{7, 5, 3}
	var biggest int
	biggest = a[0]
	for i := range a {
		if biggest < a[i] {
			biggest = a[i]
		}
	}

	fmt.Println("Biggest number is ", biggest)

}
