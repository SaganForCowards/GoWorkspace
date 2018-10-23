package main

import "fmt"

func main() {
	fmt.Print("Enter a Temperature in  Degrees Farenheit: ")
	var F float64
	fmt.Scanf("%f", &F)

	C := ((F - 32) * 5 / 9)

	fmt.Println(F, " is ", C, " in Degrees Celsius.")
}
