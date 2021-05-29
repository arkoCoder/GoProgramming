package GoRoutines

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func GetGouroutineResults() {
	fmt.Println("Starting GoRoutine examples")
	// spawns a new goroutine
	wg.Add(2)
	go sayHello()
	var msg string = "Hello"

	// Decoupling the msg variable with the goRoutine as using msg directly would cause RACE CONDITION ang would lead to improper data
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
	// Race condition
	msg = "Lag gye!!"

}

func sayHello() {
	fmt.Println("Ola!!")
	wg.Done()
}
