package main

import (
	"fmt"
	"time"
)

type nachricht int

func main() {
	fmt.Println("hello from go")
	// Deklarierung der einzelden Channals
	// Als Rohr zu verstehen
	var hinweg = make(chan nachricht)
	var rückweg = make(chan nachricht)
	var stop = make(chan nachricht)

	// Erstellung der Klassen
	// Die Channal werden ihnen umgekehrt zugewisen
	// Die Channal in den Klassen sind als anschlüsse zu verstehn an den die Rohre angelegt werden
	var Ping = Akteur{
		hinweg,
		rückweg,
	}

	// Erstellung der Klassen
	// Die Channal werden ihnen umgekehrt zugewisen
	// Die Channal in den Klassen sind als anschlüsse zu verstehn an den die Rohre angelegt werden
	var Pong = Akteur{
		rückweg,
		hinweg,
	}

	// Start der gorotinen
	go Ping.ping(stop)
	go Pong.pong(stop)

	// zum starten muss das erste Channal/Rohr gefüllt werden
	hinweg <- 0

	time.Sleep(time.Second)

}

// Deklaration der Klassen mit den Channals/ Rohranschlüssen
type Akteur struct {
	send   chan nachricht
	recive chan nachricht
}

// Funktionen ping
// func (name Klassenname)Funktionsname(Übergabeparameter)
func (a Akteur) ping(ende chan nachricht) {
	for {
		// so weißt man eine Nachricht von einem channal einer Variable zu
		zähler := <-a.send
		fmt.Println("ping : ", zähler)
		// so sendet man eine Variable über einen Channal/ Rohr
		a.recive <- zähler + 1
		if zähler >= 30 {
			ende <- 1
			return
		}

	}
}

func (a Akteur) pong(ende chan nachricht) {
	for {
		// Switch case wie in java nur ohne ()
		select {
		case stop := <-ende:
			fmt.Println("pong : Ende gelände", stop)
			return
		case zähler := <-a.send:
			fmt.Println("pong : ", zähler)
			a.recive <- zähler + 1
		}
	}
}
