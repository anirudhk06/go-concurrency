package main

import "fmt"

func main() {
	buffChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case buffChannel <- s:

		}
	}

	close(buffChannel)

	for result := range buffChannel {
		fmt.Println(result)
	}
}
