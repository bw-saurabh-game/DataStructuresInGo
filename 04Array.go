package main

import "fmt"

func Array() {

	var myArray = [5]int{1, 2, 3, 4, 5}
	myArray2 := [2]string{"sam", "saurabh"}
	arr1 := [5]int{}
	arr2 := [5]int{1: 10, 2: 40} // initializing 1 and 2 elements only

	for _, value := range myArray {
		fmt.Print(value, " ")
	}
	println()

	for _, value := range myArray2 {
		fmt.Print(value, " ")
	}
	println()

	for _, value := range arr1 {
		fmt.Print(value, " ")
	}
	println()

	for _, value := range arr2 {
		fmt.Print(value, " ")
	}
	fmt.Println()

	for i := 0; i < len(myArray); i++ {
		fmt.Print(" ", myArray[i])
	}
	fmt.Println()

	println("Length of Array can be find by :len(myArray) -->", len(myArray))
}
