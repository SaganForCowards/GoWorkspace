package main

import "fmt"

func main() {
	var numOne int
	var numTwo int

	fmt.Println("Please enter a small number: ")
	fmt.Scan(&numOne)
	fmt.Println("Please enter a large number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, " = ", numTwo%numOne)

}
