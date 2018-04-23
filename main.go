/*
 Author: Matthew Shiel

 References for the algorithm:
	* https://en.wikipedia.org/wiki/Shunting-yard_algorithm,
	* Video by lecturer Ian McLoughlin
	* https://www.youtube.com/watch?v=Wz85Hiwi5MY

 References for understanding strings and runes:
	* https://mymemorysucks.wordpress.com/2017/05/03/a-short-guide-to-mastering-strings-in-golang/
*/

package main

import (
	"fmt"

	"./shunt"
)

func main() {

	/*var infix, postfix string

	//prompt the user to enter the infix string
	fmt.Println("Enter infix expression to convert: ")
	fmt.Scanln(&infix)
	//Convert string to postfix notation
	postfix = shunt.IntToPost(infix)

	fmt.Println("Converted expression is: " + postfix)*/
	fmt.Println("Converted expression is: " + shunt.IntToPost("a.b.c*"))

}
