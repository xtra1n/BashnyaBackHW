package squares

import "sync"

func worker(num int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- num * num
}

func CalculateSquares(nums []int, result chan<- int, wg *sync.WaitGroup) {
	for _, num := range nums {
		wg.Add(1)
		go worker(num, result, wg)
	}
}

func Sum(result <-chan int) int {
	var sum int

	for num := range result {
		sum += num
	}

	return sum
}
