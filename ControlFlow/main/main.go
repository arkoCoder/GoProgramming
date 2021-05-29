package main

import (
	"main/go/ControlFlow/DeferPanicRecovery"
	"main/go/ControlFlow/IfStatements"
	"main/go/ControlFlow/Loops"
	"main/go/ControlFlow/SwitchStatements"
)

func main() {
	IfStatements.GetIfStatementResults()
	SwitchStatements.GetSwitchStatementResults()
	Loops.GetLoopsResult()
	DeferPanicRecovery.GetDPRResults()
}
