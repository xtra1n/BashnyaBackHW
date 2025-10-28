package utils

import "fmt"

func PrintError(err error) {
	fmt.Println(err)
}

func InputNumber(inputPrompt string) (int, error) {
	var num int

	fmt.Print(inputPrompt)

	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	return num, nil
}
