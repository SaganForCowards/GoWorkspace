package main

import "fmt"

func main() {
	var NumOne int
	fmt.Println("Please enter a whole number: ")
	fmt.Scan(&NumOne)
	x, y := tryharder(NumOne)
	fmt.Println(x, y)
}

func tryharder(NumOne int) (int, bool) {
	return NumOne / 2, NumOne%2 == 0
}
