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
			// 'Last In First Out' operation
			// Pop 2 things off the stack
			frag2 := nfastack[len(nfastack)-1]

			// Remove popped item and update the stack
			nfastack = nfastack[:len(nfastack)-1] // Everything on the stack up to the last character

			frag1 := nfastack[len(nfastack)-1]

			// Remove popped item and update the stack
			nfastack = nfastack[:len(nfastack)-1] // Everything on the stack up to the last character

			// Concatenate the two popped items
			frag1.accept.edge1 = frag2.initial
			//push new nfa fragment onto the stack
			nfastack = append(nfastack, &nfa{
				initial: frag1.initial, accept: frag2.accept})

		// Alternation operator
		case '|':
			// 'Last In First Out' operation
			// Pop 2 things off the stack
			frag2 := nfastack[len(nfastack)-1]

			// Remove popped item and update the stack
			nfastack = nfastack[:len(nfastack)-1] // Everything on the stack up to the last character

			frag1 := nfastack[len(nfastack)-1]

			// Remove popped item and update the stack
			nfastack = nfastack[:len(nfastack)-1] // Everything on the stack up to the last character

			// Create two new states
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}

			// Concatenate the items
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//push new nfa fragment onto the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		// Kleane star - zero or more
		case '*':
			// Pop an item off the stack
			frag := nfastack[len(nfastack)-1]

			// Remove popped item and update the stack
			nfastack = nfastack[:len(nfastack)-1]
			//point at oold accpet state & new initial state
			accept := state{}
			//new initial state , point to old initial & new accept
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial //old frag with 2 extra states
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfastack[0]

}
