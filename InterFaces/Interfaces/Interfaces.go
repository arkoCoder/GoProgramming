package Interfaces

import "fmt"

type msgWriter interface {
	WriteMsgViaInterface(msg string) int
}

type WriteMsg struct{}

func GetInterfacesResults() {
	fmt.Println("Starting Interfaces execution")

	var msg msgWriter = WriteMsg{}
	msg.WriteMsgViaInterface("Ola!")
}

// func implementing the interfaceby having the signature just like that of interface. Name of the method is same as interface
func (wrtMsg WriteMsg) WriteMsgViaInterface(msg string) int {
	n, _ := fmt.Println(msg)
	return n
}