package main

import (
	"fmt"
	"time"
)

type message int //ist nen int, heißt nur anders. kein bool, da hier auch direkt die Anzahl der Sendungen mitgezählt wird

func main() {

	var a_to_b = make(chan message) //Channel in eine Richtung
	var b_to_a = make(chan message) //Channel in andere Richtung

	var a = Actor{  //Bezeichnung, was der Aktor bei den jeweiligen Channels ist
		sender:   a_to_b,
		receiver: b_to_a,
	}
	var b = Actor{ //Bezeichnung, was der Aktor bei den jeweiligen Channels ist
		sender:   b_to_a,
		receiver: a_to_b,
	}

	//Starten der Routinen
	go a.ping()
	go b.pong()

	//Senden der ersten Nachricht
	a_to_b <- 0

	//benötigt, da Go sonst die Routinen stoppt wenn die main durch ist
	time.Sleep(time.Second)
}

type Actor struct { //Erstellung des Aktors
	receiver chan message
	sender   chan message
}

func (a Actor) ping() {
	for {
		// nur zum hin und her senden: a.sender <- a.receiver + 1
		var amount = <-a.receiver //in eigener Variable, da für
		fmt.Println("Ping: ", amount)
		if amount > 4 {
			a.sender <- -1
			return
		} else {
			a.sender <- amount + 1
		}
	}
}

func (a Actor) pong() {
	for {
		var amount = <-a.receiver
		fmt.Println("Pong:", amount)
		if amount == -1 {
			return
		}
		a.sender <- amount + 1
	}
}