package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type menuFunction struct {
	Title    string
	Function func(a *App) error
}

func generateMenu(commands map[string]menuFunction) string {
	menu := "Меню:\n"
	for i := 0; i < len(commands); i++ {
		j := strconv.Itoa(i)
		menu += fmt.Sprintf("%s - %s\n", j, commands[j].Title)
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
	return fmt.Sprintf("%d - %-20s - %s", d.Id, str, d.Date.Format("02.01.2006"))
}

func (n Note) TextToLine100(text string, width int) string {
	var result strings.Builder
	var line []rune
	words := strings.Fields(text)
	for _, word := range words {
		wordRune := []rune(word)
		if len(line)+len(wordRune)+1 <= width {
			if len(line) == 0 {
				line = wordRune
			} else {
				line = append(line, ' ')
				line = append(line, wordRune...)
			}
		} else {
			result.WriteString(string(line) + "\n")
			line = wordRune
		}
	}
	if len(line) != 0 {
		result.WriteString(string(line))
	}
	return result.String()
}

type Note struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
	Text string    `json:"text"`
}

func EnterValue(helpMsg string, resolvingVoid bool) (string, error) {
	fmt.Print(helpMsg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value := strings.TrimSpace(scanner.Text())
	if value == "" && !resolvingVoid {
		return "", errors.New("значение не может быть пустым")
	}
	return value, nil
}
