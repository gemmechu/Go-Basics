package main

import (
	"fmt"
	"math"
)

func sumDecimal(n int) {
	var a float64 = float64(1) / float64(n)
	var sum float64 = 0
	var p float64 = 0
	for i := 1; i < 1000; i++ {
		p = math.Pow(4, a)
		sum += p
	}

	fmt.Println(sum)

}
func main() {
	a := 3
	sumDecimal(a)

}
