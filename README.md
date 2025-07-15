# Go tools

A collection of tools and resources for learning using go.

Packages:
- structures: Working with queues and stacks in Go.
- help: Add utilities like Print and Sleep function


## A simple main file
```go
package main

import (
	. "github.com/school-of-decima/gotools/help"
	"github.com/school-of-decima/gotools/structures"
)

func main() {
	var mainStack = structures.NewStringStack("main")
	
	mainStack.Push("A", "B", "C")

	for mainStack.Len() > 0 {
		Print(mainStack)
		item := mainStack.Pop()
		Print(item)

		MSleep(100)
	}

	Print(mainStack)

}
```