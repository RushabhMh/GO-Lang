package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d: stopping\n", id)
            return
        default:
            fmt.Printf("Worker %d: working\n", id)
            time.Sleep(500 * time.Millisecond) // Simulate work
        }
    }
}

func main() {
    // Create a context with cancellation
    ctx, cancel := context.WithCancel(context.Background())

    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }

    // Let the workers run for a while
    time.Sleep(2 * time.Second)

    // Cancel the context to stop all workers
    fmt.Println("Main: cancelling context")
    cancel()

    // Give some time for goroutines to print their stopping message
    time.Sleep(1 * time.Second)
    fmt.Println("Main: done")
}
