package main

import (
	. "github.com/school-of-decima/gotools/help"
	"github.com/school-of-decima/gotools/structures"
)

func main() {
	var mainStack = structures.NewNamedStack[string]("main")

	mainStack.Push("A", "B", "C", "D", "E", "F", "G", "H", "I", "J")

	for mainStack.Len() > 0 {
		Print(mainStack)
		item := mainStack.Pop()
		Print(item)

		MSleep(100)
	}

	Print(mainStack)

}
