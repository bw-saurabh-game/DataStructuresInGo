package main

import "fmt"

func pointerr() {

	var digit int = 10

	pointerr := &digit
	fmt.Println(*pointerr)

	*pointerr = 100
	fmt.Println(*pointerr)

	var digit2 *int = pointerr
	fmt.Println(*digit2)
	abc := 99
	PointerPrinter(abc)

	println(abc)

}

// 	fmt.Println(&digit)

// 	fmt.Println(*pointerr)

// fmt.Println(*pointerr)

func PointerPrinter(digit int) {
	fmt.Println("Digit passed to the function:", digit)
	digit = 12
	fmt.Println("Digit after assignment", digit)
}
