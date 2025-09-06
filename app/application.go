package app

import (
	"fmt"
)

var funcTable = map[string]menuFunction{
	"0": menuFunction{Title: "Выход", Function: StopApp},
	"1": menuFunction{Title: "Посмотреть заметки", Function: ShowNotes},
	"2": menuFunction{Title: "Добавить заметку", Function: AddNotes},
}

func RunApp() {
	fmt.Println("\033[33mДобро пожаловать \"Записки Ластоногих\"\033[0m")
	var command string
	for {
		fmt.Println("----------------------------")
		fmt.Print(generateMenu(funcTable))
		fmt.Scan(&command)
		fmt.Println("----------------------------")
		targetF, ok := funcTable[command]
		if !ok {
			fmt.Println("\033[31mКоманда не найдена\033[0m")
			continue
		}
		err := targetF.Function()
		if err != nil {
			fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())

			continue
		}
		fmt.Println("\033[32mУспешно!\033[0m")
	}
}
