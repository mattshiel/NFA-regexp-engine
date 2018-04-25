package main

import (
	"fmt"

	"./nfaEngine"
	"./shunt"
)

func main() {

	var infix string
	var userInput string
	var postfix string

	// Chose to use a simple scanner instead of a buffered reader
	// If the program was expanded to where it needed to read in full sentences
	// I would swap to a reader instead

	// Print menu options
	fmt.Println("Choose an option, input your regexp, generate an NFA and match your string against it" +
		"\n1 - Enter a regular expression in infix notation" +
		"\n2 - Enter a regular expression in postfix notation" +
		"\n3 - Exit")
	fmt.Scanln(&userInput)

	for userInput != "3" {

		// Main Menu options
		switch userInput {

		// Infix to postfix
		case "1":
			// Prompt user to enter the infix regexp
			fmt.Println("Enter infix expression: ")
			fmt.Scanln(&infix)

			//Convert infix regexp to postfix
			postfix = shunt.IntToPost(infix)
			fmt.Println(postfix)

			NFAStringMatch(postfix)

		// Postfix regexp
		case "2":
			fmt.Println("Enter postfix expression: ")
			//prompt the user to enter the postfix string
			fmt.Scanln(&postfix)
			//create & test NFA
			NFAStringMatch(postfix)

		// Default case
		default:
			fmt.Println("Invalid choice please try again")
		}
		fmt.Scanln(&userInput)
	}

}

// NFAStringMatch matches and tests a user string against
// the user generated NFA
func NFAStringMatch(postfix string) {

	// User string to match against the NFA
	var userString string

	// Prompt user
	fmt.Println("Enter a string to test against the NFA: ")
	fmt.Scanln(&userString)

	// Generate the NFA and test string against it
	// then output the results
	fmt.Print("Match Result: ")
	fmt.Print(nfaEngine.PoMatch(postfix, userString))
	fmt.Println("")
}
