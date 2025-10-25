package main

import (
	"bashnya-hw2/internal/app"
	"bashnya-hw2/pkg/utils"
	"fmt"
)

func main() {
	fmt.Println("Введеное число должно быть в диапазоне (-12307; 12307)")
	num, err := utils.InputNumber()

	if err != nil {
		utils.PrintError(err)
		return
	}

	err = app.CheckNum(num)
	if err != nil {
		utils.PrintError(err)
		return
	}

	result, err := app.ProcessNumber(num)
	if err != nil {
		utils.PrintError(err)
		return
	}

	text := app.ConvertNumToText(result)

	utils.PrintResult(num, result, text)
}
