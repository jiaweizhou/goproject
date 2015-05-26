package main

import (
	//"bufio"
	"code.google.com/p/go.net/websocket"
	"fmt"
	//"io/ioutil"
	"log"
	"strings"
)

var message = make(chan []byte, 2000)

func ExampleDial() {
	origin := "http://localhost/"
	url := "ws://192.168.2.10:50000/deployment-engine/log"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	/*
		if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
			log.Fatal(err)
		}
	*/
	//var msg = make([]byte, 1024)
	go func() {
		for {
			var t = make([]byte, 4096)
			ws.Read(t)
			message <- t
		}
	}()
	fmt.Println("YYYY")
	for {
		select {
		case t := <-message:
			s := strings.TrimRight(string(t), string(0))
			s = strings.TrimRight(s, string(10))
			s = strings.TrimRight(s, string(0))
			if s != "" {
				fmt.Printf("%s", s)
			}
		}
	}
}
func main() {
	ExampleDial()
}
