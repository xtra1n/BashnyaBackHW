package pipeline

func Generate(nums []int) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for _, num := range nums {
			channel <- num
		}
	}()

	return channel
}

func Multiplier(in_channel <-chan int) <-chan int {
	out_channel := make(chan int)

	go func() {
		defer close(out_channel)
		for num := range in_channel {
			out_channel <- num * 2
		}
	}()

	return out_channel
}
