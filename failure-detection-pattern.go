package main

import (
	"fmt"
	"time"
)

// Heartbeat sends a heartbeat signal to the channel at regular intervals
func Heartbeat(heartbeatChan chan<- struct{}, interval time.Duration) {
	for {
		time.Sleep(interval)
		select {
		case heartbeatChan <- struct{}{}:
		        fmt.Println("Heartbeat sent")
		default:
			fmt.Println("Heartbeat channel is full, skipping")
		}
	}
}

// Monitor listens for heartbeat signals and detects failures
func Monitor(heartbeatChan <-chan struct{}, timeout time.Duration) {
	for {
		select {
		case <-heartbeatChan:
			fmt.Println("Heartbeat received")
		case <-time.After(timeout):
			fmt.Println("Heartbeat timeout, failure detected")
			return
		}
	}
}

func main() {
	heartbeatChan := make(chan struct{}, 1)
	heartbeatInterval := 1 * time.Second
	heartbeatTimeout := 3 * time.Second

	// Start the heartbeat goroutine
	go Heartbeat(heartbeatChan, heartbeatInterval)

	// Start the monitor goroutine
	go Monitor(heartbeatChan, heartbeatTimeout)

	// Let the program run for a while to observe the behavior
	time.Sleep(5 * time.Second)
	fmt.Println("Main: done")
}
