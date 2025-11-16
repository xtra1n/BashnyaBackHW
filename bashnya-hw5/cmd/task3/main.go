package main

import (
	concurrent_map "bashnya-hw5/internal/task3"
	"fmt"
	"sync"
)

func main() {
	safeMap := concurrent_map.NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-#%d", id)
			for j := 0; j < 10; j++ {
				safeMap.Set(key, j)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-#%d", id)
			for j := 0; j < 10; j++ {
				value, _ := safeMap.Get(key)
				_ = value
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-#%d", i)
		value, _ := safeMap.Get(key)
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Printf("Len of map: %d\n", safeMap.Len())
}
