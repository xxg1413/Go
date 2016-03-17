package main

import (
	"fmt"
)

type Interger int

type LessAdder interface {
	Less(b Interger) bool
	Add(b Interger)
}

func (a Interger) Less(b Interger) bool {

	return a < b
}

func (a *Interger) Add(b Interger) {
	*a += b
}

func main() {

	var a Interger = 1
	var b LessAdder = &a
	var d LessAdder = a

	var c Interger = 2

	fmt.Println(b.Less(c))
	fmt.Println(d.Less(c))

}
