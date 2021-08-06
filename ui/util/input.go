package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func InputString(annotation string) string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	message := "input " + annotation + "> "
	fmt.Print(message)
	for scanner.Scan() {
		input = scanner.Text()
		if input != "" {
			break
		}
		fmt.Print(message)
	}
	return strings.TrimSpace(input)
}

func InputSecretString(annotation string) string {
	var input string
	message := "input " + annotation + "> "

	for input == "" {
		fmt.Print(message)
		bytePassword, err := term.ReadPassword(int(int(os.Stdin.Fd())))
		if err != nil {
			break
		}
		input = string(bytePassword)
		fmt.Println("")
	}
	return strings.TrimSpace(input)
}

func ChoiceYN(annotation string) bool {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	message := annotation + " [y/n] > "
	fmt.Print(message)

	for scanner.Scan() {
		input = strings.TrimSpace(scanner.Text())
		if input == "y" || input == "n" {
			break
		}
		fmt.Print("please type y or n > ")
	}
	return input == "y"
}
