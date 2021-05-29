package Loops

import "fmt"

func GetLoopsResult() {
	fmt.Println("For Loop Execution")

	fmt.Println("Simple for loop execution")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Simple for loop execution with initial variable scopr to method level")
	test := 0
	for ; test < 5; test++ {
		fmt.Println(test)
	}

	fmt.Println("Simple for loop execution for doWhile interpretation")
	testDoWhile := 0
	for testDoWhile < 2 {
		fmt.Println(testDoWhile)
		testDoWhile++
	}

	fmt.Println("Simple for loop execution with Label")
BreakLabel:
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println(i * j)
			if i*j >= 2 {
				break BreakLabel
			}
		}
	}

	fmt.Println("Simple for loop execution over collections")

	NameToRelationMap := map[string]string{
		"Gopal": "Son",
		"Anil":  "Father",
	}

	for k, v := range NameToRelationMap {
		fmt.Println(k, v)
	}

	fmt.Println("Simple for loop execution over string")

	testString := "Wohoo!!"
	for k, v := range testString {
		fmt.Println(k, string(v))
	}
}
