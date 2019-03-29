package number

import (
	"fmt"
	"goplayground/number/operate"
)

// Playground playground to learn how to handle number
func Playground() {
	fmt.Println("")
	fmt.Println("--NUMBER PLAYGROUND")
	fmt.Println("")

	a := operate.Sum(4, 6)
	fmt.Println(a)
}
