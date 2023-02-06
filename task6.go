// all possible ways to stop goroutine
package main

import (
	"context"
	"fmt"
	"time"
)

func contextCancel(ctx context.Context, timeout time.Duration) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
			return
		default:
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func contextTimeout(ctx context.Context, timeout time.Duration) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context timeout")
			return
		default:
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

// cancel using cancel context
func cancelContextExample() {
	context, cancel := context.WithCancel(context.Background())
	go contextCancel(context, 5*time.Second)

	time.Sleep(6 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	/*prints:
	working...
	working...
	working...
	working...
	working...
	working...
	context canceled*/
}

// cancel using timeout context
func timeoutContextExample() {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go contextTimeout(context, 5*time.Second)

	time.Sleep(6 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	/*prints:
	working...
	working...
	working...
	working...
	working...
	context timeout*/
}

// cancel using cancel channel
func cancelChannelExample() {
	cancel := make(chan struct{})
	go func() {
		for {
			select {
			case <-cancel:
				fmt.Println("channel closed")
				return
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(6 * time.Second)
	close(cancel)
	time.Sleep(2 * time.Second)
	/*prints:
	working...
	working...
	working...
	working...
	working...
	channel closed*/
}

// cancel using timeout timeAfter
func timeoutTimeAfterExample() {
	go func() {
		timeout := time.After(2 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("time after done")
				return
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	/*prints:
	working...
	working...
	time after done*/
}

// cancel using timeout timer
func timeoutTimerExample() {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("timer timeout")
				return
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(6 * time.Second)
	/*prints:
	working...
	working...
	working...
	working...
	working...
	timer timeout*/
}

// cancel using flag
func flagCancelExample() {
	flag := false
	go func() {
		for {
			if flag {
				fmt.Println("flag canceled")
				return
			}
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(6 * time.Second)
	flag = true
	time.Sleep(2 * time.Second)
	/*prints:
	working...
	working...
	working...
	working...
	working...
	working...
	flag canceled*/
}

func main() {
	cancelContextExample()
	timeoutContextExample()
	cancelChannelExample()
	timeoutTimeAfterExample()
	timeoutTimerExample()
	flagCancelExample()
}
