package main

import "fmt"

func main() {
	i := (true && false) || (false && true) || !(false && false)
	for {
		if i == true {
			fmt.Println("True!")
		} else {
			fmt.Println("False!")
		}

	}

}

// really overthought this one.
