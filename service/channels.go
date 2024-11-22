package service

import "fmt"

func MyChannel() {
	myChannel := make(chan string)

	go func() { myChannel <- "Greeting My people! Welcome to a new world" }()

	msg := <-myChannel
	fmt.Println(msg)
}
