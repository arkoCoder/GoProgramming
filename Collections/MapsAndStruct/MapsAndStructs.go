package MapsAndStruct

import (
	"fmt"
	"reflect"
)

type Actor struct {
	name string
	age  int
}

type Bollywood struct {
	Actor
	isBollywood bool
}

type TagStruct struct {
	name string `required`
}

func GetValueOfMapsAndStruct() {
	fmt.Println("Maps and Structs")
	// ways of defining maps
	testMap := map[string]int{
		"Gopal": 25,
	}
	fmt.Printf("Test map %v %T \n", testMap, testMap)

	//using inbuilt make method
	testMap01 := make(map[int]int)
	testMap01 = map[int]int{
		1: 1,
	}
	fmt.Printf("Test map01 %v %T \n", testMap01, testMap01)

	// Adding , deleting and retrieving entries from map
	testMap01[2] = 2
	fmt.Printf("Test map01 %v %T \n", testMap01, testMap01)
	delete(testMap01, 2)
	fmt.Printf("Test map01 %v %T \n", testMap01, testMap01)
	fmt.Printf("Test map01 %v %T \n", testMap01[1], testMap01[1])

	// using ", ok" to test value exist in the map or not
	checkValue, ok := testMap01[2]
	fmt.Printf("Check map %v %v \n", checkValue, ok)

	actor := Actor{
		name: "John",
		age:  35,
	}

	embeddedStruct := Bollywood{}
	embeddedStruct.age = 60
	embeddedStruct.name = "Anil Kapoor"
	embeddedStruct.isBollywood = true

	fmt.Printf("Value for embedded Struct %v \n", embeddedStruct)

	fmt.Printf("Structure %v %T \n", actor, actor)

	embeddedStructLiteralSyntax := Bollywood{
		Actor: Actor{
			name: "Christial Bale",
			age:  50,
		},
		isBollywood: false,
	}

	fmt.Printf("Value of embeddedStructLiteralSyntax %v \n", embeddedStructLiteralSyntax)

	t := reflect.TypeOf(TagStruct{})
	field, _ := t.FieldByName("name")
	fmt.Printf("Value od the tag in struct %v %T \n", field.Tag, field.Tag)
}
