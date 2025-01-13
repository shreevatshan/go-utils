package std

import (
	"errors"
)

type Queue struct {
	Elements chan interface{}
	size     int32
}

func InitQueue(size int32) *Queue {
	queue := &Queue{
		Elements: make(chan interface{}, size),
		size:     size,
	}
	return queue
}

func (queue *Queue) Enqueue(element interface{}) error {
	select {
	case queue.Elements <- element:
		return nil
	default:
		return errors.New("queue size full")
	}
}

func (queue *Queue) Dequeue() interface{} {
	return <-queue.Elements
}

func (queue *Queue) EnqueueOrWait(element interface{}) {
	queue.Elements <- element
}

func (queue *Queue) Channel() chan interface{} {
	return queue.Elements
}
