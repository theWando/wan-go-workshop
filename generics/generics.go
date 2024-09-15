package main

import "fmt"

type Foo[T any] struct {
	Bar T
}

type number interface {
	int | float32 | float64
}

func Sum[N number](nums ...N) N {
	var n N
	for _, num := range nums {
		n += num
	}
	return n
}

func main() {
	floats := []float32{2.3, 4.5, 6.7}
	ints := []int{1, 2, 3, 4, 5, 6}

	sumIntsFn := Sum[int]
	fmt.Printf("floats %.2f\n", Sum[float32](floats...))
	fmt.Printf("ints %d\n", sumIntsFn(ints...))
}
