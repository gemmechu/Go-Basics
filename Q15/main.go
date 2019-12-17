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

func missingNumber(list []int) {
	sorted := bubbleSort(list)
	n := list[len(sorted)-1]
	result := []int{}
	for i := 1; i < n; i++ {
		found := false
		for _, j := range list {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			result = append(result, i)
		}
	}
	if len(result) == 1 {
		result = append(result, n+1)

	}
	if len(result) == 0 {
		result = append(result, n+1)
		result = append(result, n+2)
	}

	fmt.Println(result)
}

func main() {
	a := []int{1, 4, 3, 2}
	missingNumber(a)

}
