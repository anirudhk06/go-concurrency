package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	channel := make(chan int)

	go func() {

		for _, n := range nums {
			channel <- n
		}
		close(channel)
	}()

	return channel
}

func qs(channel <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for n := range channel {
			out <- n * n
		}

		close(out)

	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := qs(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
