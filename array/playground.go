package array

import "fmt"
import "goplayground/array/index"

// Playground learn how to handle array type
func Playground() {
	fmt.Println("")
	fmt.Println("--ARRAY PLAYGROUND")
	fmt.Println("")

	// 1st way create array, it will auto contain five 0 if it not replaced later: [0 0 0 0 0]
	var x [5]int
	x[4] = 89

	// 2nd way create array, remove the var, and define the values should contain it
	y := [3]int{1, 2, 3}
	y[0] = 5
	y[1] = 4

	fmt.Printf("Create array x: %v\n", x) // => [0 0 0 0 89]
	fmt.Printf("Create array y: %v\n", y) // => [5 4 3]

	a := []int{10, 11, 12}
	b := append(a, 13) // because append doesn't modify the param array, so it could abe assign as a new array variable

	fmt.Printf("Defined a: %v\n", a) // => [10, 11, 12]
	fmt.Printf("Defined b: %v\n", b) // => [10, 11, 12, 13]

	index.SetValue(a, 1, 9)          // this fn update the real array index through it's pointer, so var a become real changed
	index.SetValue(a, 2, 8)          // this fn update the real array index through it's pointer, so var a become real changed
	fmt.Printf("Updated a: %v\n", a) // => [10, 9, 8]

	newb := []int{60, 70, 80, 90}
	err := index.SetAllValue(b, newb) // tottaly update all index in b arr to a new one that given as 2nd param
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Updated b: %v\n", b) // => [60, 70, 80, 90]
	}
}
