package main

import "fmt"

func buildWord(word string, wordarray []string) {
	var count int = 0
	var minCount int = 1000
	var countlen int = 0
	for i := 0; i < len(wordarray); i++ {
		count = 0
		countlen = 0
		for j := 0; j < len(wordarray); j++ {
			if i != j {
				countlen += len(wordarray[i])
				count++
			}
			if countlen == len(word) {
				if minCount > count {
					minCount = count
				}
				break
			}
		}
	}
	fmt.Println(minCount)
}
func main() {
	buildWord("buildword", []string{"buil", "dwor", "bu", "ild", "wo", "rd"})

}
