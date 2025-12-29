package main

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

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(option), nil
}

func getOption(options []Option) (string, error) {
	keys := make([]string, len(options))
	for i, option := range options {
		keys[i] = option.key
	}
	for {
		for _, option := range options {
			fmt.Printf("%s - %s\n", option.key, option.text)
		}
		option, err := getUserInput("Choose")
		if err != nil {
			return "", err
		}
		option = strings.ToUpper(option)
		if slices.Contains(keys, option) {
			return option, nil
		}
		fmt.Printf("Invalid option '%s'\n", option)
	}
}
