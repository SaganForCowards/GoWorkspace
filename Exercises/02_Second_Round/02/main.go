package main

import "fmt"

func main() {
	NumOne := func(o int) (int, bool) {
		return o / 2, o%2 == 0

	}
	fmt.Println(NumOne(15))
}
