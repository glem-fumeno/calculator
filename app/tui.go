package app

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

type Options []Option

func NewOptions(options ...Option) Options {
	return options
}

func (o Options) Add(options ...Option) Options {
	return append(o, options...)
}

func NewOption(key, text string, a ...any) Option {
	return Option{key, fmt.Sprintf(text, a...)}
}
func NewLine(text string, a ...any) Option {
	return Option{"", fmt.Sprintf(text, a...)}
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

func GetOption(error string, options Options) string {
	keys := make([]string, 0, len(options))
	for _, option := range options {
		if option.key != "" {
			keys = append(keys, option.key)
		}
	}
	for {
		fmt.Print("\033[2J\033[H")
		fmt.Printf("\033[1;31m%s\033[0m\n", error)
		for _, option := range options {
			if option.key == "" {
				fmt.Println(option.text)
			} else {
				fmt.Printf("%s - %s\n", option.key, option.text)
			}
		}
		option := GetInput("Choose")
		option = strings.ToUpper(option)
		if slices.Contains(keys, option) {
			return option
		}
		error = fmt.Sprintf("Invalid option '%s'", option)
	}
}
