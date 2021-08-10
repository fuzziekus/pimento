package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

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

	var fd int
	if term.IsTerminal(syscall.Stdin) {
		fd = syscall.Stdin
	} else {
		bytePassword, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		input = string(bytePassword)
		return strings.TrimSpace(input)
	}

	for input == "" {
		fmt.Fprint(os.Stderr, message)
		bytePassword, err := term.ReadPassword(fd)
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
