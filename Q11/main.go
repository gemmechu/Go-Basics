package main

import "fmt"

func findExtreme(list []int) {
	min := list[0]
	max := list[0]
	for _, element := range list {
		if min > element {
			min = element
		}
		if max < element {
			max = element
		}
	}
	fmt.Println("The max: ", max)
	fmt.Println("The min: ", min)
}
func main() {
	a := []int{4, 1, 7, 8}
	findExtreme(a)

}
