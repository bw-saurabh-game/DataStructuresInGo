package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	// wg := sync.WaitGroup{}

	go abc11(30, 22, ch)
	go abc22(30, 20, ch)
	time.Sleep(11 * time.Second)

}

func abc11(val1 int, val2 int, ch chan int) {

	result := val1 + val2
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("abc11  executed ....", result)

}

func abc22(val1 int, val2 int, ch chan int) {
	result := 0
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		result += <-ch
	}

	fmt.Println("Answer of  abc22", result)
}
