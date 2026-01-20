package app

import (
	"fmt"
)

func ProcessNumber(num int) (int, error) {
	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num = num*13 + 1
			continue
		} else {
			num = (num + 2) * 3
		}

		if num%13 == 0 && num%9 == 0 {
			return 0, fmt.Errorf("service error")
		} else {
			num += 1
		}
	}

	return num, nil
}

func CheckNum(num int) error {
	var err error

	if num <= -12307 || num >= 12307 {
		err = fmt.Errorf("range error")
	}

	return err
}
