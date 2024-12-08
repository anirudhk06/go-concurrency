package main

import "fmt"

func main() {
	firstChannel := make(chan string)
	secondChannel := make(chan string)

	go func() {
		firstChannel <- "First Channel"
	}()

	go func() {
		secondChannel <- "Second Channel"
	}()

	select {
	case msgFromFirstChannel := <-firstChannel:
		fmt.Println(msgFromFirstChannel)
	case msgFromSecondChannel := <-secondChannel:
		fmt.Println(msgFromSecondChannel)
	}
	//  In this case only one channel will print out of the terminal because of the select statement

}
