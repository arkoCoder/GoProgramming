package Functions

import (
	"fmt"
)

type greet struct {
	msg string
	day string
}

func GetFuncResults() {
	fmt.Println("Get function results")
	greetingFunc("Hello", " Gopal")
	greetingFuncWithDay("Hello", "Gopal", "Monday")
	sumOfValues("Sum is: ",1,2,3)
	s := sumOfNumbers(1,2)
	fmt.Println(s)

	// Anonymous function
	func() {
		fmt.Println("Hello Anonymous Function")
	}()

	f:= func(msg string){
		fmt.Println(msg)
	}
	f("Hello function as a variable")
	
	// Methods
	g := greet{
		msg : "Hello Gopal ",
		day : "Monday",
	}
	g.greetingMethod()
}

func greetingFunc(greeting string, name string) {
	fmt.Print(greeting, name)
}

// No need to define string for each parameter. It can be defined at the end if all parameters are of same type
func greetingFuncWithDay(greeting, name , day string) {
	fmt.Printf("\n %v %v. It is %v \n", greeting, name, day)
}

//Variatic type of declaration for parameters. Such parameters should always be the last parameters in function
func sumOfValues(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result+=v
	}
	fmt.Println(msg,result)
}

// Method invocation by providing the context of struct we have created
func (g greet) greetingMethod() {
	fmt.Printf("\n %v. It is %v \n", g.msg, g.day)
}

// Way of returning named variables
func sumOfNumbers(values ...int) (result int) {
	fmt.Println(values)
	for _,v := range values {
		result+=v
	}
	return
}