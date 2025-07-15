package structures

import "fmt"

type queue[T any] struct {
	queueName string
	items     []T
}

// NewNamedQueue creates a new queue with a specified name.
func NewNamedQueue[T any](queueName string) *queue[T] {
	return &queue[T]{
		queueName: queueName,
	}
}

// NewQueue creates a new queue with a default name.
func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		queueName: "queue",
	}
}

// Add adds an item to the end of the queue.
func (q *queue[T]) Add(item ...T) {
	q.items = append(q.items, item...)
}

// Enqueue adds an item to the end of the queue.
func (q *queue[T]) Enqueue(item ...T) {
	q.Add(item...)
}

// Dequeue removes and returns the item from the front of the queue.
func (q *queue[T]) Dequeue() T {
	return q.Consume()
}

// Consume removes and returns the item from the front of the queue.
func (q *queue[T]) Consume() T {
	if len(q.items) == 0 {
		panic("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Len returns the number of items in the queue.
func (q *queue[T]) Len() int {
	return len(q.items)
}

// Print prints the queue's name and its items.
func (q *queue[T]) Print() string {
	return fmt.Sprintf("%s: %v", q.queueName, q.items)
}
