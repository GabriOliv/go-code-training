package main

import (
	"errors"
	"fmt"
)

// circularBuffer Struct
type circularBuffer struct {
	// Size of Buffer
	size int
	// Front of Queue
	front int
	// Rear of Queue
	rear int
	// previous command
	// 0 = dequeue | 1 = enqueue
	prev   int
	buffer []interface{}
}

// Constructor
func newCircularBuffer(size int) *circularBuffer {
	b := circularBuffer{size: size, buffer: make([]interface{}, size, size)}
	return &b
}

// Queue Interface
type queue interface {
	dequeue()
	enqueue()
	overwrite()
	print()
}

// Enqueue Value
func (b *circularBuffer) enqueue(input interface{}) interface{} {
	if (b.front == b.rear) && (b.prev == 1) {
		return errors.New("buffer is full")
	}

	b.buffer[b.rear] = input
	b.rear = (b.rear + 1) % b.size
	b.prev = 1

	b.print()
	return nil
}

// Dequeue Value
func (b *circularBuffer) dequeue() interface{} {
	if (b.front == b.rear) && (b.prev == 0) {
		return errors.New("buffer is empty")
	}

	output := b.buffer[b.front]
	b.buffer[b.front] = nil
	b.front = (b.front + 1) % b.size
	b.prev = 0

	b.print()
	return output
}

// Overwrite Value
func (b *circularBuffer) overwrite(input interface{}) interface{} {
	if (b.front != b.rear) || (b.prev != 1) {
		return errors.New("buffer is not full")
	}

	b.buffer[b.rear] = input
	b.rear = (b.rear + 1) % b.size
	b.front = b.rear
	b.prev = 1

	b.print()
	return nil
}

// Print Data
func (b *circularBuffer) print() {
	// cyan=same green=front red=rear reset
	color := []string{"\u001b[0m ", "\u001b[92m", "\u001b[91m", "\u001b[36m"}

	fmt.Print(" ")
	for key, slot := range b.buffer {
		switch {
		case (key == b.front) && (key == b.rear):
			fmt.Print(color[3])
		case key == b.front:
			fmt.Print(color[1])
		case key == b.rear:
			fmt.Print(color[2])
		}
		if slot == nil {
			fmt.Print("_")
		} else {
			fmt.Print(slot)
		}
		// reset
		fmt.Print(color[0])

	}

	fmt.Print("\tfront:", b.front)
	fmt.Print("\trear:", b.rear)

	switch b.prev {
	case 0:
		fmt.Print("\tprevious:", color[2], "Dequeued", color[0])
	case 1:
		fmt.Print("\tprevious:", color[1], "Enqueued", color[0])
	}

	fmt.Println()
}

func errorHandling(err interface{}) {
	switch err.(type) {
	case error:
		fmt.Println(err)
	}
}

func main() {

	// Declare a Circular Buffer
	disc := newCircularBuffer(7)
	disc.print()
	fmt.Println()

	// Move Position to 3
	for i := 1; i <= 3; i++ {
		errorHandling(disc.enqueue(i))
		errorHandling(disc.dequeue())
	}
	errorHandling(disc.dequeue())
	fmt.Println()

	// Enqueue 1 2 3
	for i := 1; i <= 3; i++ {
		errorHandling(disc.enqueue(i))
	}

	// Dequeue 1 2
	errorHandling(disc.dequeue())
	errorHandling(disc.dequeue())

	// Enqueue 4 .. 9
	for i := 4; i <= 9; i++ {
		errorHandling(disc.enqueue(i))
	}

	// Enqueue on full buffer
	errorHandling(disc.enqueue(10))
	fmt.Println()

	// Overwrite with A B
	errorHandling(disc.overwrite("A"))
	errorHandling(disc.overwrite("B"))

	// Enqueue on full buffer
	errorHandling(disc.enqueue(11))
	fmt.Println()

	// Overwrite with C D
	errorHandling(disc.overwrite("C"))
	errorHandling(disc.overwrite("D"))

}
