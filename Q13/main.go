package main

import "fmt"

func plusTwo() func(n int) int {
	return func(n int) int {
		return (n + 2)
	}
}

func main() {
	p := plusTwo()
	fmt.Println(p(2))

}
