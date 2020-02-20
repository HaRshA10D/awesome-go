package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Duration(3000000000))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tick()
		}
	}
}

func tick() {
	fmt.Printf("Time is now: %s\n", time.Now())
}
