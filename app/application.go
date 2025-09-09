package app

import (
	"fmt"
)

var funcTable = map[string]menuFunction{
	"0": menuFunction{Title: "Выход", Function: func(a *App) error { return a.StopApp() }},
	"1": menuFunction{Title: "Посмотреть заметки", Function: func(a *App) error { return a.ShowNotes() }},
	"2": menuFunction{Title: "Добавить заметку", Function: func(a *App) error { return a.AddNotes() }},
	"3": menuFunction{Title: "Посмотреть заметку подробно", Function: func(a *App) error { return a.ViewNoteFull() }},
	"4": menuFunction{Title: "Редактировать заметку", Function: func(a *App) error { return a.EditNote() }},
}

type App struct {
	Notes []Note
}

func (a *App) RunApp() {
	fmt.Println("\033[33mДобро пожаловать \"Записки Ластоногих\"\033[0m")
	var command string
	a.ReadNote()
	for {
		fmt.Println("----------------------------")
		command, _ = EnterValue(generateMenu(funcTable), false)
		fmt.Println("----------------------------")
		targetF, ok := funcTable[command]
		if !ok {
			fmt.Println("\033[31mКоманда не найдена\033[0m")
			continue
		}
		err := targetF.Function(a)
		if err != nil {
			fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())

			continue
		}
		fmt.Println("\033[32mУспешно!\033[0m")
	}
}
