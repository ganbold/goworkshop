package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working..")
	time.Sleep(5 * time.Second)
	fmt.Println("Done.")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main_21() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	messages_1 := make(chan int, 2)

	messages_1 <- 1
	messages_1 <- 2

	fmt.Println(<-messages_1)
	fmt.Println(<-messages_1)

	done := make(chan bool, 1)
	go worker(done)

	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
