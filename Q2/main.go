package main

import "fmt"

func fizzBuzz() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i%3 == 0 && i%5 == 0 {

			fmt.Println("FizzBuzz")
		} else {
			if i%3 == 0 {
				fmt.Println("Fizz")
			}
			if i%5 == 0 {
				fmt.Println("Buzz")
			}

		}
	}

}
func loopingArray() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)

}
func main() {
	fizzBuzz()
}
