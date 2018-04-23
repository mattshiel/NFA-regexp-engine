/*
 Author: Matthew Shiel

 References for the algorithm:
	* https://en.wikipedia.org/wiki/Shunting-yard_algorithm,
	* Video by lecturer Ian McLoughlin
	* https://www.youtube.com/watch?v=Wz85Hiwi5MY

 References for understanding strings and runes:
	* https://mymemorysucks.wordpress.com/2017/05/03/a-short-guide-to-mastering-strings-in-golang/
*/

/*
The Shunting Yard algorithm is stack based.
For the conversion there are two strings the input and the output.

The stack is used to hold the operators not yet added to the output queue eg. +, -, ×
The algorithm then does something depending on what operator is read in eg. "3 + 4" --> "3 4 +"

For the full algorithm in detail follow https://en.wikipedia.org/wiki/Shunting-yard_algorithm#The_algorithm_in_detail
*/

package shunt

// IntToPost converts an infix regular expression to a postfix regular expression
// and returns the resulting string to be later sent to a post fix stack evaluator
func IntToPost(infix string) string {

	// Map to store all the special characters and their precedences
	// Found regex operator precedence here https://stackoverflow.com/questions/36870168/operator-precedence-in-regular-expressions
	specials := map[rune]int{
		'*': 10,
		'.': 9,
		'|': 8,
	}

	// Array of runes to hold the postfix regular expression
	// and the stack for the infix operators
	pofix, stack := []rune{}, []rune{}

	// Loop over infix, throw the index of the char with the blank identifier as it's unneeded
	// r is the rune at the index
	for _, r := range infix {

		switch {
		// If the token is a closing bracket, push it onto the operator stack
		case r == '(':
			stack = append(stack, r)

		// If the token is a right bracket, pop the closing bracket from the stack
		// If the token on top of the operator stack is not a closing bracket pop it onto the output queue
		case r == ')':
			for stack[len(stack)-1] != '(' {
				pofix = append(pofix, stack[len(stack)-1])
				stack = stack[:len(stack)-1] // Everything on the stack up to the closing bracket (the last character)
			}

			stack = stack[:len(stack)-1] // Pop the closing bracket off the stack

		// If the rune is a special character
		// '> 0' is used because accessing a key not contained in a map returns 0
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {

				// Pop operators from the operator stack onto the output queue
				pofix = append(pofix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, r)

		// If the rune is not a bracket or a special character
		default:
			pofix = append(pofix, r)

		}

	}

	for len(stack) > 0 {
		// Pop element off of stack and onto the output queue
		pofix = append(pofix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(pofix)
}
