package main

import (
	"fmt"
)

// parameters
func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

// Here, myFunction() returns one integer (result) and one string (txt1):
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return // don't need to specify valu infront of func
}

// Recurssion
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func Function() {
	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	fmt.Println(myFunction(5, "Hello"))

	testcount(1)
}
