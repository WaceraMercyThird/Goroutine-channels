package handler

import (
	"fmt"
	"go_master/pipeline"
	"go_master/service"
	"math/rand"
	"time"
)

func SomeFuncSync() {
	//synchronous
	service.SomeFunc("1")
	service.SomeFunc("2")

	//Asynchronous( async/await
	//running Concurrency and not Parallelism

	go service.SomeFunc("3")
	go service.SomeFunc("4")

	//time duration
	time.Sleep(time.Second * 2)

	fmt.Println("Hello World")
	fmt.Println(" ")
}

func Greetings() {
	go service.Say("Hello") // Run in a goroutine

	service.Say("World")

	time.Sleep(time.Second * 2)

	fmt.Println(" ") // Run in the main thread
}

func Message() {
	service.MyChannel()

	time.Sleep(time.Second * 2)

	fmt.Println(" ")
}

func SelectedMessage() {
	service.SelectedChannel()

	time.Sleep(time.Second * 2)

	fmt.Println(" ")

}

func ForSelectedMessage() {

	//service.ForSelectChannel()
	//service.DoneChannel()

	time.Sleep(time.Second * 2)

	fmt.Println(" ")
}

func DoneChannel() {
	done := make(chan bool)

	go service.DoWork(done)

	time.Sleep(time.Second * 2)

	fmt.Println(" ")

}

//...............PIPELINES...................

func Pipeline() {
	//input
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	randNumFetchers := func() int { return rand.Intn(100) }

	pipeline.PipeLineChannel(nums)

	//PipeLine Generators / Synchronous Channels
	pipeline.PileLineGenerator()
	pipeline.PileLinePrimeGenerator(randNumFetchers, 10)
	pipeline.PileLinePrimeFanGenerator(randNumFetchers, 10)

	time.Sleep(time.Second * 2)

	fmt.Println(" ")

}

//................CHANNEL GENERATORS ..........

func Generator() {

	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int { return rand.Intn(1000000000) }

	for rando := range service.RepeatFunc(done, randNumFetcher) {
		fmt.Println("random generated numbers", rando)
	}

	time.Sleep(time.Second * 10)

	fmt.Println(" ")

}
