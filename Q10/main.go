package main

import (
	"fmt"
)
var dispatch map[string] add
func map(fun string, list) int {
	for _,i:= range list{
		add(i)
	}
}
func add(a){
	fmt.Println(a)
}

func main() {
	fib := fibonacciRecursion(9)
	fmt.Println("The fibonacci is :", fib)
}
