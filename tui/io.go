package tui

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Option struct {
	key  string
	text string
}

func NewOption(key, text string) Option {
	return Option{key, text}
}

func GetInput(prompt string) string {
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(option)
}

func GetOption(title, error string, options []Option) string {
	keys := make([]string, len(options))
	for i, option := range options {
		keys[i] = option.key
	}
	for {
		fmt.Print("\033[2J\033[H")
		fmt.Printf("\033[1;31m%s\033[0m\n", error)
		fmt.Println(title)
		for _, option := range options {
			fmt.Printf("%s - %s\n", option.key, option.text)
		}
		option := GetInput("Choose")
		option = strings.ToUpper(option)
		if slices.Contains(keys, option) {
			return option
		}
		error = fmt.Sprintf("Invalid option '%s'", option)
	}
}
