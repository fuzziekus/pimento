package ui

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"syscall"

	"github.com/fuzziekus/pimento/db"
	"golang.org/x/term"
)

type input struct {
	name       string
	annotation string
	val        string
	// isSecret   bool
}

type inputList []input

var editFlagMapping = map[string]int{
	"Description": 1,
	"UserId":      2,
	"Password":    4,
	"Memo":        8,
	"All":         15,
}

type EditFlags struct {
	All         bool
	Description bool
	UserId      bool
	Password    bool
	Memo        bool
	flagVar     int
}

func inputString(annotation string) string {
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

func inputSecretString(annotation string) string {
	var input string
	message := "input " + annotation + "> "

	for input == "" {
		fmt.Print(message)
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			break
		}
		input = string(bytePassword)
		fmt.Println("")
	}
	return strings.TrimSpace(input)
}

func (e *EditFlags) calcCondition() {
	e.flagVar = 0

	// カラム指定がある場合は、All フラグは false にし、
	// 対象カラムだけ更新するようにする
	if e.Description || e.UserId || e.Password || e.Memo {
		e.All = false
	}

	// フィールド名を動的に取得して e.flagVar を更新するため
	// reflect を利用
	rtEditFlag := reflect.TypeOf(EditFlags{})
	rvEditFlag := reflect.ValueOf(e).Elem()
	for i := 0; i < rtEditFlag.NumField(); i++ {
		f := rtEditFlag.Field(i)
		for key, val := range editFlagMapping {
			v := rvEditFlag.FieldByName(key).Interface()
			if f.Name == key && v.(bool) {
				e.flagVar |= val
			}
		}
	}
}

// 入力する変数を e.cond の値から決定し、必要なものを入力させた上で credential の obj を作成する
func (e EditFlags) GenerateCredentialFromInputInterfaces() db.Credential {
	return generateCredential(e.displayInteractiveInput())
}

func (e EditFlags) displayInteractiveInput() inputList {
	e.calcCondition()

	// 入力する変数を e.cond の値から決定する
	order := []string{"Description", "UserId", "Password", "Memo"}
	targets := []string{}
	for _, key := range order {
		if e.flagVar&editFlagMapping[key] == editFlagMapping[key] {
			targets = append(targets, key)
		}
	}

	var template inputList
	for _, s := range targets {
		tmp := input{
			name:       s,
			annotation: s + " for credential",
		}
		if s == "Password" {
			tmp.val = inputSecretString(tmp.annotation)
		} else {
			tmp.val = inputString(tmp.annotation)
		}

		template = append(template, tmp)
	}

	return template
}

func generateCredential(inputlist inputList) db.Credential {
	// inputlist の値をもとに db.Credential の値を作成
	var c db.Credential

	rtCredential := reflect.TypeOf(c)
	uv := reflect.ValueOf(&c).Elem()
	for i := 0; i < rtCredential.NumField(); i++ {
		f := rtCredential.Field(i)

		for _, input := range inputlist {
			if input.name == f.Name {
				uv.Field(i).SetString(input.val)
			}
		}
	}
	return c
}
