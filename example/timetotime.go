package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	//channel that sends a message every 100 milliseconds
	boom := time.After(500 * time.Millisecond)
	//channel that sends a message after 500 milliseconds
	for {
		select {
		case <-tick:
			//every time a message is sent to tick
			fmt.Println("tick.")
		case <-boom:
			//when a message is sent to boom
			fmt.Println("BOOM!")
			return
		default:
			//have this run every 50 milliseconds.
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
