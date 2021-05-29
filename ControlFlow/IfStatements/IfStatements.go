package IfStatements

import (
	"fmt"
)

func GetIfStatementResults() {
	fmt.Println("If statements execution")

	statePopulation := map[string]int{
		"Maharashtra": 112244,
		"Karnataka":   332233,
	}
	// .ok returns boolean whether value is present in map or not
	// initializer syntax
	if pop, ok := statePopulation["Maharashtra"]; pop > 112221 {
		fmt.Printf("Wohoo!! The population is %v and value for ok is %v \n", pop, ok)
	}

	if statePopulation["Maharashatra"] < statePopulation["Karnataka"] {
		fmt.Println("Less population is good!!")
	}
}
