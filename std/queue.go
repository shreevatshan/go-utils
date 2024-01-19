package std

import (
	"errors"
	"sync/atomic"
)

type Queue struct {
	Elements    chan interface{}
	size        int32
	currentSize atomic.Int32
}

func InitQueue(size int32) *Queue {
	queue := &Queue{
		Elements: make(chan interface{}, size),
		size:     size,
	}
	return queue
}

func (queue *Queue) getCurrentSize() int32 {
	return queue.currentSize.Load()
}

func (queue *Queue) incrementQueueSize() {
	queue.currentSize.Add(1)
}

func (queue *Queue) decrementQueueSize() {
	queue.currentSize.Add(-1)
}

func (queue *Queue) Enqueue(element interface{}) error {
	var err error
	if queue.getCurrentSize() < queue.size {
		queue.incrementQueueSize()
		queue.Elements <- element
	} else {
		err = errors.New("queue size full")
	}
	return err
}

func (queue *Queue) Dequeue() interface{} {
	var element interface{}
	if queue.getCurrentSize() > 0 {
		element = <-queue.Elements
		queue.decrementQueueSize()
	} else {
		element = nil
	}
	return element
}

func (queue *Queue) EnqueueOrWait(element interface{}) {
	queue.Elements <- element
	queue.incrementQueueSize()
}

func (queue *Queue) DequeueOrWait() interface{} {
	element := <-queue.Elements
	queue.decrementQueueSize()
	return element
}
