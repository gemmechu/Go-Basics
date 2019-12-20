package main

import (
	"fmt"
)

var start int = 0
var end int = 0

func reverseParentheses(s string) string {

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			start = i
		}
		if s[i] == ')' {
			end = i

			return reverseParentheses(s[:start] + s[start+1 : end][:len(s)] + s[end+1:])
		}

	}
	return s

}
func main() {
	a := "foo(bar)baz"
	b := reverseParentheses(a)
	fmt.Printf(b)
}
