package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// wg := sync.WaitGroup{}

	go abc3(10, 20, ch)
	go abc2(30, 20, ch)

	time.Sleep(4* time.Second)

}

func abc3(val1 int, val2 int, ch chan int) {

	fmt.Println("Started abc3 function")
	result := val1 + val2
	// time.Sleep(2*time.Second)
	ch <- result
	fmt.Println("Answer of both the values in abc3", result)
	fmt.Println("Started abc3 function done ....")

}

func abc2(val1 int, val2 int, ch chan int) {
	abd := <-ch
	fmt.Println("Started abc2 function")
	time.Sleep(2*time.Second)
	result := val1 + val2

	result += abd
	fmt.Println("Answer of both the values in abc2", result)

	fmt.Println("Started abc2 function done ....")
}
