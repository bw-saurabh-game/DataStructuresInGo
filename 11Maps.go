// A map is an unordered and changeable collection that does not allow duplicates.
// The length of a map is the number of its elements. You can find it using the len() function.

// Maps are references to hash tables.
// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

package main

import (
	"fmt"
)

func Map() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}
