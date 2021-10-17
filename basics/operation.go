package main

import "fmt"

// START OMIT
type Number struct {
	n int
}

func (n *Number) Sum(number int) {
	n.n += number
	fmt.Printf("Sum: %d\n", n.n)
}

func (n Number) Get() int {
	return n.n
}

// END OMIT

func main() {
	n := Number{6}
	n.Sum(4)
	fmt.Printf("main: %d\n", n.Get())
}
