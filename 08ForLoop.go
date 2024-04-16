package main

import "fmt"

func ForLoop() {

	//Go has only for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	var arr = []int{1, 2, 3, 4, 5, 6}

	//use of range
	for iad, ele := range arr {
		fmt.Print(iad, " ", ele, "  ")
	}

	fmt.Println()
	i := 0
	// Loop continues as long as i is less than 5
	for i < 5 {
		fmt.Print(i, " ")
		i++ // Increment i to avoid infinite loop
	}
}
