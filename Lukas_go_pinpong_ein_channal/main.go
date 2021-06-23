package main

import (
	"fmt"
	"time"
)

func PingPong(channel chan string, send string) {
	for {
		select {
		case ball := <-channel:
			fmt.Println(ball)
		case channel <- send:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func ping(channel chan string) {
	for {
		channel <- "ping..."
		ball := <- channel
		fmt.Println(ball)
	}
}

func pong(channel chan string) {
	for {
		ball := <- channel
		if (ball == "stop") {
			fmt.Println(ball)
			break
		} else {
			fmt.Println(ball)
		}
		channel <- "pong..."
	}
}

func main() {
	fmt.Println("Lets go")
	channel := make(chan string)

	//go PingPong(channel, "ping...")
	//go PingPong(channel, "pong...")

	go ping(channel)
	go pong(channel)

	time.Sleep(50 * time.Millisecond)
	channel <- "stop"
	time.Sleep(50 * time.Millisecond)

}
