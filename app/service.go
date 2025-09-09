package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func (a *App) StopApp() error {
	fmt.Println("\033[33mДо свидания!\033[0m")
	os.Exit(0)
	return nil
}

func Read() ([]Note, error) {
	f, err := os.OpenFile("data.json", os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return []Note{}, fmt.Errorf("файл не открыт: %s", err.Error())
	}
	var byteData []byte
	var items []Note
	byteData, err = io.ReadAll(f)
	if err != nil {
		return []Note{}, fmt.Errorf("ошибка: %s", err.Error())
	}
	err = json.Unmarshal(byteData, &items)
	if err != nil {
		return []Note{}, fmt.Errorf("данные не взяты: %s", err.Error())
	}
	defer f.Close()
	return items, nil
}

func (a *App) ReadNote() ([]Note, error) {
	data, err := Read()
	if err != nil {
		return []Note{}, err
	}
	a.Notes = data
	return data, nil
}

func (a *App) ShowNotes() error {
	var str []string
	for i := range a.Notes {
		str = append(str, a.Notes[i].String())
	}
	fmt.Println(strings.Join(str, "\n"))
	return nil
}

func (a *App) Safe() error {
	jsonNote, err := json.Marshal(a.Notes)
	if err != nil {
		return fmt.Errorf("данные не в json: %s", err.Error())
	}
	err = os.WriteFile("data.json", jsonNote, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("данные не в файле: %s", err.Error())
	}
	return nil
}

func (a *App) AddNotes() error {
	var newId int
	for _, note := range a.Notes {
		newId = max(newId, note.Id)
	}
	newDate := time.Now()
	newDateStr := newDate.Format("02.01.2006")
	newName, err := EnterValue("Введи название замтки: ", true)
	if err != nil {
		return err
	}
	newText, err := EnterValue("Введи описание замтки: ", true)
	if err != nil {
		return err
	}
	newNote := Note{Id: newId + 1, Name: newName, Date: newDateStr, Text: newText}
	a.Notes = append(a.Notes, newNote)
	err = a.Safe()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) IdNoteFull() (int, error) {
	id, err := EnterValue("Введи номер заметки: ", false)
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	chsNote := a.Notes[result-1]
	fmt.Printf("%d - %-20s - %s\n%s\n", chsNote.Id, chsNote.Name, chsNote.Date, chsNote.txt(chsNote.Text, 100))
	return result, err
}

func (a *App) ViewNoteFull() error {
	_, err := a.IdNoteFull()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) EditNote() error {
	pass, err := EnterValue("\nВведи пароль: ", true)
	if err != nil {
		return err
	}
	if pass != "123" {
		fmt.Println("Не правильно")
		return err
	}
	id, err := a.IdNoteFull()
	if err != nil {
		return err
	}
	name, err := EnterValue("\nВведи название заметки: ", true)
	if err != nil {
		return err
	}
	if name != "" {
		a.Notes[id-1].Name = name
		fmt.Println("Заголовок изменен!")
	} else {
		fmt.Println("Заголовок не изменен!")
	}
	text, err := EnterValue("Введи текст заметки: ", true)
	if err != nil {
		return err
	}
	if text != "" {
		a.Notes[id-1].Text = text
		fmt.Println("Текст изменен!")
	} else {
		fmt.Println("Текст не изменен!")
	}
	date, err := EnterValue("Введите дату создания заметки: ", true)
	if err != nil {
		return err
	}
	if date != "" {
		a.Notes[id-1].Date = date
		fmt.Println("Дата создания изменена!")
	} else {
		fmt.Println("Дата создания не изменена!")
	}
	err = a.Safe()
	if err != nil {
		return err
	}
	return nil
}
