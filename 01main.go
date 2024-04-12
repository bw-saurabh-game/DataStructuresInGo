package main

import "fmt"

func main() {

	fmt.Println(" ----------01 Variables function is called------------")
	var variable int
	variable = 100
	varr := "Hii I am saurabh"
	fmt.Println(variable)
	fmt.Println(varr)

	const abc int = 10
	fmt.Println("Const variable is : ", abc)

	dynamicDeclaration := 99
	fmt.Println("Dynamic variable variable is : ", dynamicDeclaration)
	var z, y = 10, "Saurabh"
	fmt.Println("variable without giving type", z, y)

	var ab, cd, me, ze int = 1, 2, 3, 4
	fmt.Println("Multiple variable declaration in single line", ab, cd, me, ze)

	fmt.Println(" --------------02 output function is called ---------------")
	output()

	fmt.Println(" ----------03 BasicDataTypes function is called------------")
	DataTypes()

	fmt.Println(" ----------04 Array function is called------------")
	Array()

}
