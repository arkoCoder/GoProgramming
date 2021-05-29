package Channels

import (
 "fmt"
 "sync"
)

var wg = sync.WaitGroup{}
func GetChannelsResults() {
	fmt.Println("Starting Channel examples")
	fmt.Println("Basic example of channel")

	ch := make(chan int)
	go func() {
		i := <- ch
		fmt.Println(i)
	}()

	go func() {
		i:= 42
		ch <- i
	}()


}
