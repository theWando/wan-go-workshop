package main

import "fmt"

// START OMIT
type Wolf struct {
}

func (w Wolf) Howl()  {
	fmt.Println("Auuuuuuuuu!")
}

type Dog struct {
	Wolf
}

func (d Dog) Bark() {
	fmt.Println("Goof!")
}

func canHowl(_ Wolf) bool {
	return true
}
// END OMIT

func main() {
	d := Dog{}
	d.Howl()
	d.Bark()

	canHowl(d.Wolf)
}
