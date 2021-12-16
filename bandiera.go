package main

import "fmt"

func main() {
	n := 9

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				fmt.Print("O")
			} else if j > i {
				fmt.Print("L")
			} else {
				fmt.Print("M")
			}
		}
		fmt.Println()
	}

}
