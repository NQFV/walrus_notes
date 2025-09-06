package app

import (
	"errors"
	"fmt"
)

type menuFunction struct {
	Title    string
	Function func() error
}

func generateMenu(commands map[string]menuFunction) string {
	menu := "Меню:\n"
	for k, v := range commands {
		menu += fmt.Sprintf("%s - %s\n", k, v.Title)
	}
	menu += "Введите номер команды: "
	return menu
}

func (d Note) String() string {
	strRune := []rune(d.Name)
	if len(strRune) > 20 {
		strRune = append(strRune[0:17], '.', '.', '.')
	}
	str := string(strRune)
	return fmt.Sprintf("%d - %-20s - %s", d.Id, str, d.Date)
}

type Note struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

func EnterValue(helpMsg string) (string, error) {
	fmt.Print(helpMsg)
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//value := strings.TrimSpace(scanner.Text())
	var value string
	fmt.Scan(&value)
	if value == "" {
		return "", errors.New("значение не может быть пустым")
	}
	return value, nil
}
