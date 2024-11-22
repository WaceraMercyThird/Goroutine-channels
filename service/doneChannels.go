package service

import (
	"fmt"
	"time"
)

func DoWork(done <-chan bool) {

	//Donechannel
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

	time.Sleep(time.Second * 1)
}
