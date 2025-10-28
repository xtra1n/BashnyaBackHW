package main

import (
	bst "bashnya-hw3"
	"bashnya-hw3/pkg/utils"
	"errors"
	"fmt"
)

func main() {
	tree := bst.New[int]()

	count, err := utils.InputNumber("Введите количество элементов дерева:\n> ")

	if err != nil {
		fmt.Errorf("service error")
		utils.PrintError(err)
		return
	}

	fmt.Print("Введите элементы дерева:\n")

	for i := 0; i < count; i++ {
		value, err := utils.InputNumber("> ")

		if err != nil {
			fmt.Errorf("service error")
			utils.PrintError(err)
			return
		}

		tree.Insert(value)
	}

	fmt.Println("Дерево успешно заполнено\n")

	searchingValue, err := utils.InputNumber("Введите значение для поиска:\n> ")

	if err != nil {
		errors.New("ошибка ввода")
		utils.PrintError(err)
		return
	}

	isSearched := tree.Search(searchingValue)

	if isSearched {
		fmt.Println("Введенное значение найдено\n")
	} else {
		fmt.Println("Введенное значение не найдено\n")
	}

	var valueToDelete int

	valueToDelete, err = utils.InputNumber("Введите значение для удаления:\n> ")

	if err != nil {
		errors.New("ошибка ввода")
		utils.PrintError(err)
		return
	}

	isDeleted := tree.Remove(valueToDelete)

	if isDeleted {
		fmt.Println("Введенное значение удалено\n")
	} else {
		fmt.Println("Введенное значение не удалено")
	}
	treeDepth := tree.Depth()

	fmt.Printf("Глубина дерева: %d\n", treeDepth)

}
