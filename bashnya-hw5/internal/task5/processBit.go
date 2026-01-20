package processBit

func SetBit(num int64, index int, flag bool) int64 {
	var result int64

	if flag {
		result = num | (1 << index)
	} else {
		result = num &^ (1 << index)
	}

	return result
}
