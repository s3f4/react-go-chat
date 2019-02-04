package main

import (
	"fmt"
	"math"
)

// test using a function
// its naber
type test func(int, int) int

type A interface {
	test()
}

// Ax formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
type Ax struct {
	testx test
}

func main() {
	fmt.Println(math.Acos(23.3))
	fmt.Println("naber")
	a := Ax{}
	a.testx = func(a int, b int) int {
		return a + b
	}
	fmt.Print(a.testx(1, 2))

}
