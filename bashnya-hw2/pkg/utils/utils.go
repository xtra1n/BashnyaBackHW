package utils

import "fmt"

func InputNumber() (int, error) {
	var num int

	fmt.Print("Введите целое число: ")

	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	return num, nil
}

func PrintResult(original, result int, text string) {
	fmt.Printf("Введенное число: %d\n", original)
	fmt.Printf("Итоговое число: %d: %s\n", result, text)
}

func PrintError(err error) {
	fmt.Println(err)
}
