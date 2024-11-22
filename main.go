package main

import (
	"fmt"
	"go_master/service"
	"time"
)

func main() {
	service.SomeFunc("1")
	service.SomeFunc("2")
	go service.SomeFunc("3")
	go service.SomeFunc("4")

	time.Sleep(time.Duration(7000))

	fmt.Println("Hello World")
}
