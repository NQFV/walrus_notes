package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func StopApp() error {
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

func ShowNotes() error {
	notes, err := Read()
	if err != nil {
		return err
	}
	var str []string
	for _, note := range notes {
		str = append(str, note.String())
	}
	fmt.Println(strings.Join(str, "\n"))
	return nil
}

func Write(newNote Note) error {
	notes, err := Read()
	if err != nil {
		return err
	}
	notes = append(notes, newNote)
	jsonNote, err := json.Marshal(notes)
	if err != nil {
		return fmt.Errorf("данные не в json: %s", err.Error())
	}
	err = os.WriteFile("data.json", jsonNote, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("данные не в файле: %s", err.Error())
	}
	return nil
}

func AddNotes() error {
	var newId int
	notes, err := Read()
	if err != nil {
		return err
	}
	for _, note := range notes {
		newId = max(newId, note.Id)
	}
	newDate := time.Now()
	newDateStr := newDate.Format("02.01.2006")
	newName, err := EnterValue("Введи название замтки: ")
	if err != nil {
		return err
	}
	newNote := Note{Id: newId + 1, Name: newName, Date: newDateStr}
	err = Write(newNote)
	if err != nil {
		return err
	}
	return nil
}
