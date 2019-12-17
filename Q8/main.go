package main

import (
	"fmt"
)

type Channel struct {
	value int
}

func vararg() {
	var i int
	fmt.Printf("To stop enter -1")
	var channels []Channel
	for true {
		fmt.Scanf("%d", &i)
		if i == -1 {
			break
		}
		channels = append(channels, Channel{value: i})
	}
	for _, element := range channels {
		fmt.Print(element.value)
	}

}

func main() {
	vararg()
}
