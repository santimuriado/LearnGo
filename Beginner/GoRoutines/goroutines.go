package main

import (
	"fmt"
	"sync"
)

//We can remove the time.Sleep with this.
var wg = sync.WaitGroup{}

/*
func main() {

	msg := "Hi"
	go func() {
		fmt.Println(msg)
	}()
	// Go is waiting for the time.Sleep line so that it can execute the Println so it prints "Bye"
	// when
	//there is no waitgroup.
	msg = "Bye"
	//time.Sleep generally should not be used. It ties execution to dead time.
	time.Sleep(100 * time.Millisecond)
}
*/
// With go run -race the go compiler provides extra info on memory problems dealing with
// concurrency.VERY USEFUL.

func main() {

	msg := "Hi"
	//Adds a goroutine to the waitgroup (wg).
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		//Tells the wg that the goroutine ends.
		wg.Done()
	}(msg)
	msg = "Bye"
	//Tells the wg that it has to synchronize.
	wg.Wait()

}
