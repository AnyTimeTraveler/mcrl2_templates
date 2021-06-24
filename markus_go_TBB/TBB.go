package main

import "fmt"

type Bit bool

func main() {

	var com = make(chan Bit)
	var input = make(chan Bit)
	var output = make(chan Bit)
	var tbb = TBB{
		OBB{
			send:   com,
			recive: input,
			buffer: false,
		},
		OBB{
			send:   output,
			recive: com,
			buffer: false,
		},
	}

	tbb.run()

	input <- true
	fmt.Println(<-output)

	input <- true
	fmt.Println(<-output)

}

type OBB struct {
	send   chan Bit
	recive chan Bit
	buffer Bit
}

func (a OBB) empfangen() {

	a.buffer = <-a.recive
	fmt.Println("recive: ", a.buffer)
}

func (a OBB) senden() {

	fmt.Println("send: ", a.buffer)
	a.send <- a.buffer
}

type TBB struct {
	left  OBB
	right OBB
}

func (a OBB)run() {
	for {
		a.empfangen()
		a.senden()
	}
}

func (a TBB)run()   {
	go a.left.run()
	go a.right.run()
}