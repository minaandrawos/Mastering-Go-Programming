package main

import (
	"fmt"
	"time"
)

func main() {
	tickC := make(chan *time.Ticker)
	done := make(chan bool)
	go tickCounter(1, tickC, done)
	ticker := <-tickC
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(15 * time.Second)
	fmt.Println("Exiting main ..")
}

func tickCounter(n int, tickC chan *time.Ticker, done chan bool) {
	ticker := time.NewTicker(time.Duration(n) * time.Second)
	tickC <- ticker
	i := 0
Loop:
	for {
		select {
		case t := <-ticker.C:
			i++
			fmt.Println("Count", i, "at", t)
		case <-done:
			fmt.Println("done signal")
			break Loop
		}
	}

	fmt.Println("Exiting tick counter ...")
}
