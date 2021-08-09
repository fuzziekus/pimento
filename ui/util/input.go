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
	fmt.Fprint(os.Stderr, message)
	for scanner.Scan() {
		input = scanner.Text()
		if input != "" {
			break
		}
		fmt.Fprint(os.Stderr, message)
	}
	return strings.TrimSpace(input)
}

func InputSecretString(annotation string) string {
	var input string
	message := "input " + annotation + "> "

	for input == "" {
		fmt.Fprint(os.Stderr, message)
		bytePassword, err := term.ReadPassword(int(int(os.Stdin.Fd())))
		if err != nil {
			break
		}
		input = string(bytePassword)
		fmt.Fprintln(os.Stderr, "")
	}
	return strings.TrimSpace(input)
}

func ChoiceYN(annotation string) bool {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	message := annotation + " [y/n] > "
	fmt.Fprint(os.Stderr, message)

	for scanner.Scan() {
		input = strings.TrimSpace(scanner.Text())
		if input == "y" || input == "n" {
			break
		}
		fmt.Fprint(os.Stderr, "please type y or n >")
	}
	return input == "y"
}
