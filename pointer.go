package main

import "fmt"

func zeroval(tayyar int) {
	tayyar = 5
	fmt.Println("Writinf from function value ", tayyar)
}

func zeropointer(tayyar *int) {
	*tayyar = 30
}

func main() {
	i := 3
	zeroval(i)
	fmt.Println("First iteration after zeroval", i)
	zeropointer(&i)
	fmt.Println("Second iteration after zeroval", i)
}
