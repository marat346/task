package main

import "fmt"

func main() {
	a := 5
	b := 2
	fmt.Println(a, b)
	fmt.Print("----------------/n")
	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)

}
