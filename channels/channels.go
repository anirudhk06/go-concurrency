package main

import (
	"fmt"
	"time"
)

func main() {

	myChan := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		myChan <- "Data"
	}()

	msg := <-myChan

	fmt.Println(msg)

}
