package main

import (
	"fmt"
	"time"
)

func greet(phrase string, ch chan bool) {
	fmt.Println("Hello!", phrase)
	ch <- true
}

// func slowGreet(phrase string) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// }

// func main() {
// 	go greet("Nice to meet you!")
// 	go greet("How are you?")
// 	go slowGreet("How ... are ... you ...?")
// 	go greet("I hope you're liking the course!")
// }


func slowGreet(phrase string,  ch chan bool) {

	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	 ch <- true  // Send a signal to indicate the completion of the task
    close(ch)   // Close the channel after sending the signal
	// The for range ch loop terminates automatically when the channel is closed and all values are received.
}

func main() {
	// This channel is used to signal when a goroutine has finished its work.
	 ch := make(chan bool) // Create a channel for synchronization

    // Start multiple goroutines
	go greet("Nice to meet you!", ch)
	go greet("How are you?", ch)
	go slowGreet("How ... are ... you ...?", ch)
	go greet("I hope you're liking the course!",ch)
	 // Wait for the slow greeting to finish
    // <-ch // Block until signal received
	
	// Block until the slow greeting (or any task) sends a signal to the channel
    for range ch {} // The main goroutine waits here for signals
	// The loop will stop once the channel is closed and all values have been received.
}