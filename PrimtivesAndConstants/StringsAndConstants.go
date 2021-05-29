package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	//iota is a special keyword used for Enumerated constant. iota is always initializes to 0 and is incremented for subsequent values
	isHuman = 2 << iota
	isBeast
)

// Global level variable and can be used in different package as well
var I string = "Anil"

// Package level lower case variable and visibility is detained to this package only
var i string = "Indra"

func main() {
	var roles byte = isHuman | isBeast
	fmt.Printf("%b %T \n", roles, roles)
	fmt.Printf("%b %T \n", isHuman, isHuman)
	fmt.Printf("%b %T \n", isBeast, isBeast)
	fmt.Println(I, i)
	// Shadowing concept as this variable is overwriting the value of package level variable
	i = "Gopal"
	// Block level visibility for such variables as they are defined under a block
	var j string = "Harsh"
	k := 43
	fmt.Println("Test new code")
	log.Println("Logging")
	fmt.Println(I, i, j, k)
	var isBool bool = true
	fmt.Printf("%v %T \n", isBool, isBool)
	a := 1 == 1
	fmt.Printf("%v %T \n", a, a)
	var b int
	fmt.Printf("%v %T \n", b, b)

	const cst int = 10
	fmt.Println(cst)

	apiPath := "api/1.0/devices/test-device-id/enrollmentStatus"
	strArr := strings.Split(apiPath, "/")
	fmt.Println(strArr[0])
}
