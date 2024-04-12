package main

import (
	"fmt"
)

func output() {

	// Types of output function we can use
	var i, j string = "Hello", "World"
	fmt.Print(i, j) // prints with default format
	fmt.Print(i, j) //inserts space if nethier are strings
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	//The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:

	var l, m string = "saurabh", "game"
	fmt.Println(l, m)

	//The Printf() function first formats its argument based on the given formatting verb and then prints them.
	// 	%v is used to print the value of the arguments
	// %T is used to print the type of the arguments

	//     %d: Signed integer (decimal).
	// %f: Floating-point number (decimal notation).
	// %s: String.
	// %t: Boolean (true or false).
	// %v: Default format. It formats the value according to its type.
	// %x, %X: Hexadecimal representation (lowercase or uppercase).
	// %o: Octal representation.
	// %b: Binary representation.
	// %c: Unicode character.
	// %p: Pointer (base 16).
	// %e, %E: Scientific notation (lowercase or uppercase).
	// %g, %G: General format. It chooses %f or %e/%E` based on the value.

	var x, y string = "sam", "Sir"
	fmt.Printf("This is integer: %d  This is string: %s second string : %s", 100, x, y)
	fmt.Printf("x has value: %v and type: %T", x, x)
	fmt.Println()

}
