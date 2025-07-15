package structures

func NewStringStack(name string) *stack[string] {
	return NewNamedStack[string](name)
}

func NewIntStack(name string) *stack[int] {
	return NewNamedStack[int](name)
}

func NewStringQueue(name string) *queue[string] {
	return NewNamedQueue[string](name)
}

func NewIntQueue(name string) *queue[int] {
	return NewNamedQueue[int](name)
}
