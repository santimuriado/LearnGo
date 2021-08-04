package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

//A struct without fields in Golang doesn't require memory allocation.
//It works as a SIGNAL-OLNY channel.

func main() {
	go logger()
	defer func() {
		close(logCh)
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is opening"}

	logCh <- logEntry{time.Now(), logInfo, "App is closing"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

/*
func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02 T:15:04:05"), entry.severity, entry.message)
	}
}
*/

// It's very common to monitor channels with a signal only in this case for the logger with the
// doneCH.
func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf(
				"%v - [%v]%v\n",
				entry.time.Format("2006-01-02 T:15:04:05"),
				entry.severity,
				entry.message,
			)
		case <-doneCh:
		}
	}
}
