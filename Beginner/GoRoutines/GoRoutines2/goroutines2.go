package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

/*Any thread can read the data but only one at a time can write.
  Also if any thread is reading no thread can write. */
var m = sync.RWMutex{}

func main() {

	/* This sets the number of threads to be used. As minimum it's better to set it to the number of
	   threads your machine has. Best practice is to raise it but being careful because too many
	   may cause you to experience memory overhead and scheduler problems, which can make your program slower. */
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		//Locks before the goroutine is executed and closes it inside the function.
		m.RLock()
		/*Without the lock and the unlock the sayHi routine would print the numbers in a random order
		  as the threads are running against each other. */
		go sayHi()
		m.Lock()
		go incrementar()
	}
	wg.Wait()

	/*With the -1 value it shows you the amount of threads available.
	  fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))*/

}

func sayHi() {
	fmt.Printf("Hola #%v\n", counter)
	m.RUnlock() //ReadUnlock
	wg.Done()
}

func incrementar() {
	counter++
	m.Unlock()
	wg.Done()
}

/* THINGS TO CONSIDER
   -Don't build goroutines in libraries. Better for the consumer of the library to controle concurrency himself.
   -Check for race conditions con "go run -race <filename>"
   -When creating goroutines be mindful to when it ends to avoid memory leaks.
*/
