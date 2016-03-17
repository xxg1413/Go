package main

import (
	"fmt"
)

func print_info(slice []int) {

	fmt.Printf("slice: %d cap=%d len=%d\n", slice, cap(slice), len(slice))
}

func main() {

	slice1 := make([]int, 5, 10)

	slice1 = []int{0, 0, 0, 0, 0}

	print_info(slice1)

	slice1 = append(slice1, 1, 2, 3)
	slice1 = append(slice1, slice1...)

	print_info(slice1)
}
