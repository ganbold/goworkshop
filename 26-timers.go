package main

import (
	"fmt"
	"time"
)

func main_26() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C

	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
		fmt.Println("Timer 2 fired")
	}

	time.Sleep(2 * time.Second)
}
