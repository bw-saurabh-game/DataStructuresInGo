package main

import "fmt"

type ArthmaticOperation interface {
	Add() int
	Subtract() int
	Multiplication() int
	Division() int
}

type ThreeNums struct {
	a int
	b int
	c int
}

func (ThreeNums ThreeNums) Add() int {

	return ThreeNums.a + ThreeNums.b + ThreeNums.c

}

type TwoNums struct {
	a int
	b int
}

func (TwoNums TwoNums) Add() int {

	return TwoNums.a + TwoNums.b

}

func main() {

	x := TwoNums{a: 10, b: 20}
	fmt.Println("Sum of two numbers: ", x.Add())

	ThreeNum := ThreeNums{a: 10, b: 20, c: 30}

	fmt.Println("Sum of three number :", ThreeNum.Add())

}
