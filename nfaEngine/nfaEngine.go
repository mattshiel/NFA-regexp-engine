/*
 Author: Matthew Shiel

 References for Thomson's construction in C and Golang:
	* https://swtch.com/~rsc/regexp/regexp1.html,
	* Video by lecturer Ian McLoughlin
	* https://www.youtube.com/watch?v=Wz85Hiwi5MY

 Refernces used for further understanding of the history behind the algorithm
	* https://en.wikipedia.org/wiki/Thompson%27s_construction#Rules
*/

package nfaEngine

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

// poregtonfa converts a postfix regexp to an NFA and returns a pointer to it
func poregtonfa(pofix string) *nfa {

	// Stack structure
	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		// Concatenation operator
		case '.':

		// Alternation operator
		case '|':

		// Kleane star - zero or more
		case '*':
		}
	}

	return nfastack[0]

}
