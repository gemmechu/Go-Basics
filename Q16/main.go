package main

import "fmt"
var reverse = (str) => str.split('').reverse().join('');

var reverseParentheses(s string) string {
    while (s.includes('(')) {
        var l = s.lastIndexOf('(');
        var r = s.indexOf(')', s.lastIndexOf('('));
        s := s.[slice(0, l)] + reverse(s.slice(l + 1, r)) + (r + 1 === s.length ? s.slice(r, -1) : s.slice(r + 1));
    }
    return s;
}

func main() {
	a := []int{1, 4, 3, 2}
	missingNumber(a)

}
