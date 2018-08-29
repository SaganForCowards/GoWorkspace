package main

import "fmt"

func main() {
	n := largest(5, 6, 78, 2004, 62, 4057)
	fmt.Println(n)
}

func largest(jk ...int) int {
	var max int
	for _, v := range jk {
		if v > max {
			max = v
		}
	}
	return max
}
