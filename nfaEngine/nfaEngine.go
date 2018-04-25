/*
 Author: Matthew Shiel

 References for Thomson's construction in C and Golang:
	* https://swtch.com/~rsc/regexp/regexp1.html,
	* Video by lecturer Ian McLoughlin

 Refernces used for further understanding of the history behind the algorithm
	* https://en.wikipedia.org/wiki/Thompson%27s_construction#Rules
*/

package nfaEngine

import "fmt"

// State struct with a rune symbol and edges that have pointers to other states (like a linked list)
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
			// Push nfa fragment onto the stack
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
			//point at old accpet state & new initial state
			accept := state{}
			//new initial state, point to old initial & new accept
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		// One or more
		case '+':
			// pop a single element off the stack
			frag := nfastack[len(nfastack)-1]

			// Remove popped item from stack
			nfastack = nfastack[:len(nfastack)-1]

			// Accept state
			accept := state{}
			// New initial state
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})

		// Zero or more
		case '?':
			// pop a single element off the stack
			frag := nfastack[len(nfastack)-1]

			// Remove popped item from stack
			nfastack = nfastack[:len(nfastack)-1]

			// state pointing to popped item and accept state
			initial := state{edge1: frag.initial, edge2: frag.accept}
			// add the nfa to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfastack[0]

}

// PoMatch processes a string through an NFA
// and returns a boolean if it matches or not
func PoMatch(po string, s string) bool {

	// Evalutation for matching
	ismatch := false

	// Create automata
	if len(po) < 1 {
		fmt.Println("Error Invalid Expression!")

	} else {
		ponfa := poregtonfa(po)

		// Initial state
		current := []*state{}
		// Tracks what states the program can get to from current
		next := []*state{}

		current = addState(current[:], ponfa.initial, ponfa.accept)

		// Loop over s character by character
		for _, r := range s {
			// Current state
			for _, c := range current {
				// Check symbols against eachother
				if c.symbol == r {
					//
					next = addState(next[:], c.edge1, ponfa.accept)
				}
			}
			current, next = next, []*state{}
		}

		// Loop through current state array
		for _, c := range current {
			// Check for accepts
			if c == ponfa.accept {
				ismatch = true
				break
			}
		}
	}
	return ismatch
}

// addState is used to add a state to an NFA struct
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)
	// empty string arrows
	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}
