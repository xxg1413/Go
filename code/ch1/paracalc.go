package main

import "fmt"

func sum(values []int, resultChan chan int) {
	sum := 0
	for i, value := range values {
		fmt.Println("i=", i+1)
		sum += value
	}

	resultChan <- sum
}

func main() {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultChan := make(chan int, 2)

	go sum(values[:9], resultChan)
	go sum(values[9:], resultChan)

	sum1, sum2 := <-resultChan, <-resultChan

	fmt.Println("RESULT:", sum1, sum2, sum1+sum2)
}
