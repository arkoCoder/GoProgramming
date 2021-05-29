package DeferPanicRecovery

import (
	"fmt"
	"log"
)

func GetDPRResults() {
	fmt.Println("Defer Panic And Recovery results")

	fmt.Println("Defer basic example")

	fmt.Println("start")
	//As it is defer this will execute once method execution is completed
	defer fmt.Println("middle")
	fmt.Println("end")

	// Defer sets the value of the variable prior to when it is called
	a := 1
	defer fmt.Println(a)
	a = 2

	fmt.Println("Panic basic examples")

	fmt.Println("Woah All is Good Nigga")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
	panic("Oops!! something went bad")
	fmt.Println("try to recover")
}
