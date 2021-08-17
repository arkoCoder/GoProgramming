package ArraysAndSlices

import (
	"GoProgramming/Collections/MapsAndStruct"
	"fmt"
)

func GetArraysAndSlices() {
	fmt.Println("Arrays")
	// defining an array of size 4 to have values of type int with only one value to initialize with
	fixedSizeArrayWithValue := [4]int{21}
	fmt.Println(fixedSizeArrayWithValue)

	// defining an array with default value
	var fixedSizeArrayWithDefaultValue [2]int
	fmt.Printf("testArray %v \n", fixedSizeArrayWithDefaultValue)

	// defining an array of any size
	anySizeArray := [...]int{1}
	fmt.Println(anySizeArray)

	// defining arrays of array
	// [size of array][value count that can be occupied in the value array]
	var arrayOfArray [1][2]int
	arrayOfArray[0] = [2]int{1, 2}
	fmt.Println(arrayOfArray)

	// array copy example
	a := [...]int{1, 2}
	var b [2]int = a
	// value of a will be unaffected as b is completely pointing to different memory location as of a
	b[1] = 0
	fmt.Println(a)
	fmt.Println(b)

	// value of a will be affected as c is pointing to same memory location as of a
	c := &a
	c[1] = 3
	fmt.Println(a)
	fmt.Println(c)

	fmt.Println("Slices")

	testSlice01 := []int{7, 8}
	fmt.Println(testSlice01)

	var testSlice []int
	fmt.Println(testSlice)
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))
	testSlice = append(testSlice, 1)
	fmt.Println(testSlice)
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))

	// this will affect testSlice01 as both slices are pointing to same array
	testSlice02 := testSlice01
	testSlice02[1] = 2
	fmt.Println(testSlice01)
	fmt.Println(testSlice02)

	// different ways of creating slices
	testSlice03 := []int{1, 2, 3, 4, 5}
	testSlice04 := testSlice03[:]
	testSlice05 := testSlice03[2:]
	testSlice06 := testSlice03[:3]
	testSlice07 := testSlice03[2:4]
	testSlice08 := make([]int, 3, 4) // creating a slice of length 3 and capacity 4
	testSlice08[1] = 32
	testSlice09 := append(testSlice08, 1, 3, 6, 7)
	fmt.Println(testSlice09)
	fmt.Println(testSlice08)
	fmt.Println(testSlice04)
	fmt.Println(testSlice05)
	fmt.Println(testSlice06)
	fmt.Println(testSlice07)
	MapsAndStruct.GetValueOfMapsAndStruct()
	// Retrieving and deleting value from struct
	//responseFromUtils := Utils.ServerKeyFromStuct().IdmServerKey
	//fmt.Printf("Value from ssm %v %T \n", responseFromUtils, responseFromUtils)
}
