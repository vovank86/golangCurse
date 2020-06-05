package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, val := range slice {
		if val%2 != 0 {
			fmt.Println(i, "Odd")
		} else {
			fmt.Println(i, "Even")
		}
	}
}
