package main

import (
	"fmt"
)

// Converts an infix regexp to and returns a postfix regexp
func Intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	pofix := []rune{}
	s := []rune{}
	return string(pofix)
}

func main() {

	fmt.Println("Infix:     ", "a.b.c")
	fmt.Println("Postfix:     ", Intopost("a.b.c"))

}
