package main

import (
	"fmt"
	"unicode/utf8"
)

func buildTriangle() {
	for i := 1; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}
func countString() {
	var temp string = "asSASA ddd dsjk"
	count := utf8.RuneCountInString(temp)
	fmt.Println("count:", count)
}
func reverse() {
	temp := "foobar"
	runes := make([]rune, utf8.RuneCountInString(temp))
	i := len(runes)
	for _, rune := range temp {
		i--
		runes[i] = rune
	}
	fmt.Println(string(runes[i:]))
}
func main() {
	reverse()
}
