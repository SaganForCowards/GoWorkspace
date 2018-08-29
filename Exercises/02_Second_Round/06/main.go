package main

import "fmt"

func main() {
	h := 0
	for i, j := 0, 1; j < 4000001; i, j = i+j, i {
		if i%2 == 0 {
			h = +i
		}
	}
	fmt.Println(h)
}
