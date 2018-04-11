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

package main

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

	return string(pofix)
}

func main() {

}
