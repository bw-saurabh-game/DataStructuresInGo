package main

import "fmt"

func DataTypes() {
	var a int = 10
	var b bool = false
	var c string = "Saurabh"
	var d float32 = 10.99
	var x uint = 500 //can only store non-negative values:

	fmt.Println(" Integer is :", a)
	fmt.Println(" Boolean is :", b)
	fmt.Println(" String is :", c)
	fmt.Println(" float is :", d)
	fmt.Println(" uint is :", x)

	// ways to declare variables
	var a0 int
	var a1 int = 10
	var a2, a3 int = 10, 20
	var a4 = 100
	a5 := 200

	fmt.Println(a0, a1, a3, a4, a5, a2)

}
