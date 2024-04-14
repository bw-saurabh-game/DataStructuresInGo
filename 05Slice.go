package main

import "fmt"

func Slice() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// slice from array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// how to create slices using the make() function:
	myslice12 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice12)
	fmt.Printf("length = %d\n", len(myslice12))
	fmt.Printf("capacity = %d\n", cap(myslice12))

	// with omitted capacity
	myslice22 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice22)
	fmt.Printf("length = %d\n", len(myslice22))
	fmt.Printf("capacity = %d\n", cap(myslice22))

	// changing element of slice
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//append
	prices = append(prices, 90, 60)
	fmt.Println("After appending numbers prices are :", prices)

	//append slices
	myslice3 := append(prices, myslice12...)
	fmt.Println("After appending numbers prices & myslice12 slices :", myslice3)

	//Change The Length of a Slice
	arr12 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice13 := arr12[1:5]                // Slice array
	fmt.Printf("myslice1 = %v\n", myslice13)
	fmt.Printf("length = %d\n", len(myslice13))
	fmt.Printf("capacity = %d\n", cap(myslice13))

	myslice13 = arr12[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice13)
	fmt.Printf("length = %d\n", len(myslice13))
	fmt.Printf("capacity = %d\n", cap(myslice13))

	myslice13 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice13)
	fmt.Printf("length = %d\n", len(myslice13))
	fmt.Printf("capacity = %d\n", cap(myslice13))

	//The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
