# NFA-regex-engine

## Author: Matthew Shiel 

## Student ID: G00338622

A program which can build a NFA from a regexp and test a given string against the NFA to test for a match. This program was developed as a project for my third year 'Graph Theory' module in GMIT. 

## Original Problem Statement 
> You must write a program in the Go programming language [2] that can build a non-deterministic finite automaton (NFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text. You must write the program from scratch and cannot use the regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “∗” have the special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1∗ means any number of 1’s. These special characters must be used in your submission.

## How to Run The NFA Engine

*Assumes that Git and Go are installed along with the prerequisites.*
**If not, they can be found from https://golang.org/dl/ and https://git-scm.com/downloads**

**1. Clone the Repository**
```bash
> git clone https://github.com/mattshiel/NFA-regexp-engine.git
```
**2. Change Directory to the Folder**

```bash
Open the terminal/command line and navigate into the folder 
eg. > cd NFA-regexp-engine
```

**2. Compile the NFA Engine**

```bash
> go build main.go
```

**3. Run the NFA Engine**

```bash
To run the program enter './' followed by the executable produced
For Mac:
> ./main
For Windows:
> ./main.go.exe

Alternatively:
> go run main.go
```

## Program Guide
1. Select whether to enter in a regular expression in infix or postfix notation.
2. Enter in the regular expression.
3. Enter in the string to match against the generated NFA.
4. Receive output -> press enter to repeat the process.
5. 'Hold ctrl/control' + 'z' to force quit the program.

## Basics of the Program
The engine can convert infix notation to postfix notation. It does this through a simple implementation of the shunting yard algorithm. Next, a small series of nfa fragments are created from the regular expression. These fragments make up the final NFA. This is then passed a string to match against. If the string matches against the regular expression then the output is true, otherwise the output is false.

## Design Choices

### Shunting Yard Algorithm
To solve the issue of converting infix to postfix notation I used the shunting yard algoritm.
This algorithm respects the operator precedences too.

For the conversion there are two strings, the input and the output.
A stack structure is used to hold the operators not yet added to the output queue eg. +, -, ×, ?
I used Map to store all the special characters and their precedences.
I then implemented the rules for the algorithm.

These can be found here: https://en.wikipedia.org/wiki/Shunting-yard_algorithm

## Thomson's Construction
I followed Thomson's construction to help build the NFA. As learnt from Dr. Cox's article on this, NFAs, DFAs and regular expressions all have the same computational power. Unlike Dr. Cox's implementation in C however I did not try to add a DFA implementation and simply stuck to the NFAs. 

> The NFA for a regular expression is built up from partial NFAs for each subexpression with a different construction for each operator.

ie. The overall NFA is made up from lots of smaller NFA fragments.
These smaller fragments are stored in a stack structure in my program.

The overall NFA in my program is then reflected as linked collection of state structures. This NFA can then be used to test a string against.

## Extras

1. The user has menu options for adding an infix or postfix notation regexp.
2. The '?' operator was added as an extra to the project.
3. The '+' operator was added as an extra to the project.
4. The NFA engine can recognize strings over the alphabet [a-z] and [A-Z]

## References

* [Ian McLoughlin](https://github.com/ianmcloughlin) *lecturer of Graph Theory at G.M.I.T*
* Shunting yard algorithm :	https://en.wikipedia.org/wiki/Shunting-yard_algorithm,
* Russ Cox - Regular Expression Matching Can Be Simple And Fast 
(but is slow in Java, Perl, PHP, Python, Ruby, ...) : https://swtch.com/~rsc/regexp/regexp1.html
* Introuction To The Theory Of computation : http://fuuu.be/polytech/INFOF408/Introduction-To-The-Theory-Of-Computation-Michael-Sipser.pdf
* Please note any adapted code snippets have been referenced in the code.

I consulted the Golang docs, Russ Cox article and videos by lecturer Ian McLoughlin frequently over the course of my assignment.
