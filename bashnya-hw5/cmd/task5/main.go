package main

import (
	processBit "bashnya-hw5/internal/task5"
	"fmt"
)

func main() {
	var num int64 = 33
	var index int = 4
	var modified_num int64

	modified_num = processBit.SetBit(num, index, true)
	fmt.Printf("for num = %d, index of bit changed to 1 = %d: new num = %d\n", num, index, modified_num)

	modified_num = processBit.SetBit(num, index, false)
	fmt.Printf("for num = %d, index of bit changed to 0 = %d: new num = %d\n", num, index, modified_num)
}
