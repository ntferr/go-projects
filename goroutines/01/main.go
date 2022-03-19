package main

import (
	"log"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		log.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
