package main

import (
	pipeline "bashnya-hw5/internal/task4"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	num_channel := pipeline.Generate(nums)

	multiplied_channel := pipeline.Multiplier(num_channel)

	for i := range multiplied_channel {
		fmt.Printf("multiplied nums: %v\n", i)
	}
}
