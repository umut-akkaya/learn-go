package main

import (
	"fmt"
)

type Person struct {
	name    string
	surname string
	age     int
	hp      int
	str     int
	intel   int
	sta     int
}

func main() {
	ali := Person{"Ali", "Veli", 23, 100, 10, 6, 7}
	descriptionTemplate := `{"name":"%s","surname":"%s","health":"%d","age":"%d"}`
	fmt.Println(fmt.Sprintf(descriptionTemplate, ali.name, ali.surname, ali.hp, ali.age))
}
