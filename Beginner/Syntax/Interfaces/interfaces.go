package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	//Polymorphic behaviour.
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello world"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCLoser()

	//Converts the string into a byte slice.
	wc.Write([]byte("Hello everyone"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	//Useful way to check if the conversion is done properlly.
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Failed")
	}
}

//Interfaces describe behaviour not data storage.
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {

	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCLoser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//Another example.
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

/*It's impossible to create methods for primitive types as int,string etc...
  What is possible is to create a custom type that's built with the primitives and make methods
  of those. */

/*It's always better to use multiple smaller interfaces than one big interface.
  Also to not export interface which aren't necessary in other files. */
