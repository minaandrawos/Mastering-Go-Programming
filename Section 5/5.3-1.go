package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := <-tickCounter(1)
	time.Sleep(10 * time.Second)

	ticker.Stop()

	time.Sleep(15 * time.Second)

	fmt.Println("Exiting..")
}

func tickCounter(n int) <-chan *time.Ticker {

	tickC := make(chan *time.Ticker)
	go func() {
		ticker := time.NewTicker(time.Duration(n) * time.Second)
		tickC <- ticker
		i := 0
		for t := range ticker.C {
			i++
			fmt.Println("Count", i, "at", t)
		}
	}()
	return tickC
}
