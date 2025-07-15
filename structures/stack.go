package structures

import "fmt"

type stack[T any] struct {
	stackName string
	items     []T
}

// NewStack creates a new stack.
func NewNamedStack[T any](stackName string) *stack[T] {
	return &stack[T]{
		stackName: stackName,
	}
}

// NewStack creates a new stack.
func NewStack[T any]() *stack[T] {
	return &stack[T]{
		stackName: "stack",
	}
}

// Push adds an item to the top of the stack.
func (s *stack[T]) Push(item ...T) {
	s.items = append(s.items, item...)
}

// Pop removes and returns the item from the top of the stack.
func (s *stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Len returns the number of items in the stack
func (s *stack[T]) Len() int {
	return len(s.items)
}

// Print prints the stack's name and its items
func (s *stack[T]) Print() string {
	return fmt.Sprintf("%s: %v", s.stackName, s.items)
}
