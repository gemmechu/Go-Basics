package main

import "fmt"

func bubbleSort(numbers []int) []int {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	return numbers

}
func main() {
	a := []int{4, 1, 7, 8}
	b := bubbleSort(a)
	fmt.Println(b)

}
