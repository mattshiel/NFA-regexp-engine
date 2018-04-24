package main

import (
	"fmt"

	"./nfaEngine"
	"./shunt"
)

func main() {

	var infix string
	var userInput string

	// Chose to use a simple scanner instead of a buffered reader
	// If the program was expanded to where it needed to read in full sentences
	// I would swap to a reader instead

	// Print menu options
	fmt.Println("1 - Enter a regular expression in infix notation" + "\n2 - Exit")
	fmt.Scanln(&userInput)

	for userInput != "2" {

		// Main Menu options
		switch userInput {
		case "1":
			// Prompt user to enter the infix regexp
			fmt.Println("Enter infix expression to generate NFA: ")
			fmt.Scanln(&infix)

			//Convert infix regexp to postfix
			postfix := shunt.IntToPost(infix)

			NFAStringMatch(postfix)

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
