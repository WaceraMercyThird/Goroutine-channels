package pipeline

import (
	"fmt"
	"go_master/service"
	"math/rand"
	"runtime"
	"time"
)

func PipeLineChannel(nums []int) {
	//stage1
	dataChannel := service.SliceToChannel(nums)
	//stage
	finalChannel := service.Sq(dataChannel)
	// stage3
	for n := range finalChannel {
		fmt.Println("Data Received", n)

	}

}

func PileLineGenerator() {
	// Signal channel
	done := make(chan int)

	// Close the channel when function completes
	defer close(done)

	// Generate random numbers
	randNumFetcher := func() int { return rand.Intn(50000) }
	randIntStream := service.RepeatFunc(done, randNumFetcher)
	// Use pipeline functions
	for rando := range service.Take(done, randIntStream, 50) {
		fmt.Println("Random generated number:", rando)
	}

	time.Sleep(time.Second * 2)

	fmt.Println(" ")

}

func PileLinePrimeGenerator(randNumFetcher func() int, count int) {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	// Use pipeline functions
	randIntStream := service.RepeatFunc(done, randNumFetcher)
	primeStream := service.PrimeFinder(done, randIntStream)
	for rando := range service.Take(done, primeStream, count) {
		fmt.Println("Random generated prime number:", rando)
	}

	fmt.Println("Time Duration taken to complete task:", time.Since(start))
}

func PileLinePrimeFanGenerator(randNumFetcher func() int, count int) {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	// Use pipeline functions
	randIntStream := service.RepeatFunc(done, randNumFetcher)
	//primeStream := PrimeFinder(done, randIntStream)

	//fan out
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = service.PrimeFinder(done, randIntStream)
	}

	//fan in
	fannedInStream := service.FanIn(done, primeFinderChannels...)

	for rando := range service.Take(done, fannedInStream, count) {
		fmt.Println("Random generated prime number:", rando)
	}

	fmt.Println("Time Duration taken to complete task:", time.Since(start))
}
