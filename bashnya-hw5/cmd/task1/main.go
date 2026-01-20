package main

import (
	squares "bashnya-hw5/internal/task1"
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	result := make(chan int, len(nums))

	squares.CalculateSquares(nums, result, &wg)

	go func() {
		wg.Wait()
		close(result)
	}()

	sum := squares.Sum(result)

	fmt.Printf("Sum = %d\n", sum)
}
