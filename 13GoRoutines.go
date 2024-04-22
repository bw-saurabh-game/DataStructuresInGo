package main

import (
	"fmt"
	"sync"
	// "time"
)

func GoRoutines() {

	fmt.Println("Main Function got start")

	wg := &sync.WaitGroup{}

	count := 10

	for i := 0; i < 10; i++ {

		wg.Add(2)
		count = 9
		go abc(wg, &count)
		// wg.Wait()
		// wg.Add(1)
		go abcd(wg, &count)
		// count = 9
		// fmt.Println("count is  :", count)
		wg.Wait()
	}

	fmt.Println("Main Function got executed")
}

var mutex sync.Mutex

func abc(wg *sync.WaitGroup, count *int) {
	defer wg.Done()
	fmt.Println("function started.....")
	fmt.Println("count is in abc :", *count)
	mutex.Lock()
	*count = 100
	mutex.Unlock()
	// time.Sleep(2 * time.Second)
	fmt.Println("function completed.....")

}

func abcd(wg *sync.WaitGroup, count *int) {
	defer wg.Done()
	fmt.Println("abcd started.....")
	fmt.Println("count is in abcd :", *count)
	mutex.Lock()
	*count = 10
	mutex.Unlock()
	// time.Sleep(3 * time.Second)
	fmt.Println("abcd completed.....")

}
