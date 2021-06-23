package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan PingPongBall) //Channel erstellen

	//Funktionen starten
	go ping(messages)
	go pong(messages)

	time.Sleep(time.Second)
}

type PingPongBall struct{ //aka Klasse
	//um Funktionen der Klasse zuzuweisen: func (NameInFunktion KlassenName) funktionsname(Übergabeparameter)
	text string
	zaehler int
}
func ping(channel chan PingPongBall) { //Akteur mit Datentyp: Name(Übergabename Datentyp ÜbergabeDatentyp)
	zaehler:=0 //erstellen der Zählervariable
	inhalt := "ping"
	for {
		//wenn anzahl von PingPongs größer ist als 10, dann wird ein stop gesendet
		if (zaehler >= 10) {
			inhalt = "stop"
		}
		channel <- PingPongBall{inhalt, zaehler + 1} //Nachricht an den Channel senden
		messageReceived := <-channel //Nachricht aus dem Channel erhalten
		fmt.Println("Ping" , messageReceived) //erhaltene Nachricht ausgeben
		zaehler = messageReceived.zaehler //zähler aus der erhaltenen Nachricht übernehmen
	}
}

func pong(channel chan PingPongBall) {
	for { //zum unendlichen laufen
		messageReceived := <-channel
		fmt.Println("Pong", messageReceived)
		//Abbruch, wenn ein stop empfangen wird
		if(messageReceived.text == "stop") {
			break
		}
		//neuer Ball zurück mit dem Zähler aus der Originalnachricht
		channel <- PingPongBall{"pong", messageReceived.zaehler + 1}
	}
}

