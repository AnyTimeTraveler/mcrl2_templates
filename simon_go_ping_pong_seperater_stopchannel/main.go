package main

import (
	"fmt"
	"time"
)

type message int

type Actor struct {
	mesgChannel chan message
	stopChannel chan string
}

func dummy(a ...interface{}) {

}

func main() {
	var mesgChannel = make(chan message)
	var stopChannel = make(chan string)

	var a = Actor{mesgChannel, stopChannel}
	var b = Actor{mesgChannel, stopChannel}

	dummy(a, b, mesgChannel, stopChannel)

	go a.ping()
	mesgChannel <- 0
	go b.pong()

	time.Sleep(time.Millisecond)
	stopChannel <- "stop it now!"

	time.Sleep(time.Second)
	test()
}

func (a Actor) ping() {
	for {
		select {
		case amount := <-a.mesgChannel:
			fmt.Println("Ping:", amount)
			a.mesgChannel <- amount + 1
		case stopMsg := <- a.stopChannel:
			fmt.Println("Ping:", stopMsg)
			return
		}
	}
}

func (a Actor) pong() {
	for {
		var amount = <-a.mesgChannel
		fmt.Println("Pong:", amount)
		a.mesgChannel <- amount + 1
	}
}

func test()  {
	var x = '0'

	var i = 0

	i = int(x)
	fmt.Println(i)
}
