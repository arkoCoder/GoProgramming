package SwitchStatements

import (
	"fmt"
)

func GetSwitchStatementResults() {
	fmt.Println("Switch statements execution")

	switch 1 {
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}

	switch i := 2 + 4; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd number")
	case 2, 4, 6, 8:
		fmt.Println("Even Number")
	}

	name := ""
	switch {
	case name == "Gopal":
		fmt.Println("This is Gopal")
	case name == "Harsh":
		fmt.Println("This is Harsh")
	default:
		fmt.Println("A guy has no name")

	}

	//type switches
	var testType interface{} = "1"
	switch testType.(type) {
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("Other Type")
	}
}
