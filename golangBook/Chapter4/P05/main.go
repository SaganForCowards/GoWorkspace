package main

import "fmt"

func main() {
	fmt.Print("Enter a Length in feet: ")
	var F float64
	fmt.Scanf("%f", &F)

	M := (F * 0.3048)

	fmt.Println(F, " feet is equal to ", M, " Meters.")
}
