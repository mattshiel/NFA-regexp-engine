package main

import (
	"fmt"

	"./shunt"
)

func main() {

	var infix, postfix string

	//prompt the user to enter the infix string
	fmt.Println("Enter infix expression to convert: ")
	fmt.Scanln(&infix)
	//Convert string to postfix notation
	postfix = shunt.IntToPost(infix)

	fmt.Println("Converted expression is: " + postfix)
}
