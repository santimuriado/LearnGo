package main

import (
	"fmt"
	"sync"
)

// Channels help to transfer data between goroutines in a safe way and avoids race conditions and
// memory sharing problems.

var wg = sync.WaitGroup{}

func main() {

	//Creates an int channel. Channels are strongly typed, only ints can go through.
	ch := make(chan int)
	wg.Add(2)
	//Recieving go routine
	go func() {
		//Data from the channel into the variable i. The arrow indicates data flow.
		i := <-ch
		fmt.Println(i)
		ch <- 42
		wg.Done()
	}()
	/*This is an example of a deadlock.
	  It's when each member of a group waits for another to take action, so they both do nothing.
	  Sending go routine
	*/
	/*
		go func() {
			i := 23
			ch <- i
			i = 27
			wg.Done()
		}()
	*/
	/*This is how it should work.
	  Sending go routine.
	*/
	go func() {
		ch <- 23
		fmt.Println(<-ch)
		wg.Done()
	}()
	// In this case both functions work as readers and writers but it's better if each function does
	// only one job.
	wg.Wait()

}
