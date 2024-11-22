package service

import (
	"fmt"
	"time"
)

func Say(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}
