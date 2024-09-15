package main

import "fmt"

// An interface representing only the type int.
type A interface {
	int
}

// An interface representing all types with underlying type int.
type B interface {
	~int
}

// An interface representing all types with underlying type int which implement the String method.
type C interface {
	~int
	String() string
}

// An interface representing an empty type set: there is no type that is both an int and a string.
type D interface {
	int
	string
}

type IP interface {
	string | []byte
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *leaf[T]
}

type leaf[T any] struct {
	val         T
	left, right *leaf[T]
}

func (bt *Tree[T]) find(val T) **leaf[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).val); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

type InnerInt int

func fooInts[T ~int](t T) {
	fmt.Printf("%v", t)
}
func main() {
	ii := InnerInt(2)
	fooInts(ii)

	fooInts(3)
}
