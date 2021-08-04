package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	/* With the 50 parameter it makes a buffered channel which solves the panic but it loses the second value
	   in this case.
	   One of the best purposes of a buffered channel is to work with recievers and senders that work at different frequencies.
	   If the sender has to send a lot of data in burst transmissions, the reciever is going to have problems handling
	   so much data so to handle it better it needs the buffered channel. This way it could still work while processing
	   the data.
	*/
	ch := make(chan int, 50)
	wg.Add(2)
	//Solo recieving go routine.
	/*go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		Cant send data into the channel because it's a recieve only channel.
		//ch <- 42.
		wg.Done()
	}(ch)
	*/
	// Prints both values in the for loop but goes into deadlock because it doesn't know how to exit
	// the loop.
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	/*
		go func(ch <-chan int) {
			for {
				if i, ok := <-ch; ok {
					fmt.Println(i)
				} else {
					break
				}
			}
			wg.Done()
		}(ch)
	*/
	//Solo sending go routine.
	go func(ch chan<- int) {
		ch <- 23
		ch <- 15
		close(ch)
		/* Closing the channel solves the deadlock mentioned in line 30.
		   There is no primitive in go to know if a channel is closed but to make it panic.
		   Best solution would be with a defer function and a recover so that it doesn't panic.
		*/
		wg.Done()
	}(ch)
	// Problem to solve here it's that I send 2 values through the channel but i'm only recieving 1
	// in the function in line 22.
	//It only prints 23 but not 15, it goes into DEADLOCK.
	wg.Wait()
}
