package help

import (
	"fmt"
	"strings"
)

type Printable interface {
	Print() string
}

// Print prints the items passed to it.
func Print(item ...any) {
	finalItems := make([]string, 0, len(item))
	for _, i := range item {
		switch v := i.(type) {
		case Printable:
			finalItems = append(finalItems, v.Print())
		default:
			finalItems = append(finalItems, fmt.Sprintf("%v", v))
		}
	}
	fmt.Printf("%v\n", strings.Join(finalItems, " "))
}
