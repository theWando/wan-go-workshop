package main

import "fmt"

type Number struct {
	n int
}

// START OMIT
func (n Number) Sum(number int) {
	n.n += number
	fmt.Printf("Sum: %d\n", n.n)
}

func main() {
	n := Number{6}
	n.Sum(4)
	fmt.Printf("main: %d\n", n.Get())
}

// END OMIT

func (n Number) Get() int {
	return n.n
}
