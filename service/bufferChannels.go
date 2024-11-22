package service

import (
	"fmt"
)

// For-Select Loop
func ForSelectChannel() {
	//async/await
	//buffer channel we use cue like functionality by send-and-forget or fire-and-forget
	charChannel := make(chan string, 4)
	chars := []string{"a", "b", "c", "d"}

	for _, char := range chars {
		select {
		case charChannel <- char:
		}
	}

	close(charChannel)

	for char := range charChannel {
		fmt.Println(char)
	}

}
