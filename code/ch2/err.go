package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) (sum int, err error) {

	if a < 0 || b < 0 {
		err = errors.New("more than zero")
		return
	}

	return a + b, nil
}

func main() {

	sum, err := Add(-1, 2)
	fmt.Printf("%d + %d = %d  %s  \n", 1, 2, sum, err)

}
