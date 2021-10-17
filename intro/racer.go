package main

import (
	"time"
)

// START OMIT
var count int

func race() {
	count++
}

func main() {
	go race()
	go race()
	time.Sleep(1 * time.Second)
}

// END OMIT
