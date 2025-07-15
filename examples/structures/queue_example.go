package main

import (
	"github.com/school-of-decima/gotools/help"
	"github.com/school-of-decima/gotools/structures"
)

func main() {
	queue := structures.NewNamedQueue[string]("main")

	queue.Enqueue("A", "B", "C")

	for queue.Len() > 0 {
		item := queue.Consume()

		help.Print(item)
		help.Print(queue)
	}
}
