package main

import "go_master/handler"

func main() {
	handler.SomeFuncSync()
	handler.Greetings()
	handler.Message()
	handler.SelectedMessage()
	handler.ForSelectedMessage()
	handler.DoneChannel()
	//handler.PipelineChannel()

	//goroutine concept and patterns / Synchronous Channels
	handler.Generator()
	handler.Pipeline()
}
