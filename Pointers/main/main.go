package main

import (
	"fmt"
	"main/go/Pointers/Pointers"
)

type TestPointer struct {
	foo int
}

func main() {
	Pointers.GetPointersResults()
	fmt.Println("Executing simple pointers")

	var a int = 23
	// Here b is pointing to address where value for a is stored
	var b *int = &a
	// Now if * is added before b this would be de-referencing the value for the address b is pointing to
	// Here both a and *b are pointing to same underlying data!
	fmt.Println(a, *b)

	fmt.Println("Implementing pointer for structures")
	var pointerTest *TestPointer
	pointerTest = &TestPointer{
		foo: 1,
	}
	fmt.Println(*pointerTest)

	fmt.Println("Implementing pointer using new")
	var pointerNew *TestPointer
	//New keyword can not be used with object initialization syntax
	pointerNew = new(TestPointer)
	fmt.Println(pointerNew)
}
