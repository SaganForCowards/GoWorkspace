package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%f", &input)

	if input%2 == 0 {
		fmt.Println(input, "true")
	} else {
		fmt.Println(input, "false")
	}

}
