package main

import "fmt"

func ascending(a int, b int) {
	if a > b {
		fmt.Println(b, ",", a)
	} else {
		fmt.Println(a, ",", b)
	}

}
func main() {
	ascending(6, 8)

}
