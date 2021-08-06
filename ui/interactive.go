package ui

import (
	"log"
	"reflect"

	"github.com/fuzziekus/pimento/config"
	"github.com/fuzziekus/pimento/db"
	"github.com/fuzziekus/pimento/ui/util"
)

type input struct {
	name       string
	annotation string
	val        string
}

type inputList []input

func (e *Flags) calcCondition() {
	e.flagVar = 0

	// カラム指定がある場合は、All フラグは false にし、
	// 対象カラムだけ更新するようにする
	if e.Description || e.UserId || e.Password || e.Memo {
		e.All = false
		e.NoPass = false
	}

	// フィールド名を動的に取得して e.flagVar を更新するため
	// reflect を利用
	rtEditFlag := reflect.TypeOf(Flags{})
	rvEditFlag := reflect.ValueOf(e).Elem()
	for i := 0; i < rtEditFlag.NumField(); i++ {
		f := rtEditFlag.Field(i)
		for key, val := range FlagMapping {
			v := rvEditFlag.FieldByName(key).Interface()
			if f.Name == key && v.(bool) {
				e.flagVar |= val
			}
		}
	}
}

// 入力する変数を e.cond の値から決定し、必要なものを入力させた上で credential の obj を作成する
func (e Flags) GenerateCredentialFromInputInterfaces() db.Credential {
	return generateCredential(e.displayInteractiveInput())
}

func (e Flags) displayInteractiveInput() inputList {
	e.calcCondition()

	// 入力する変数を e.cond の値から決定する
	targets := []string{}
	for _, key := range ColumnOrder {
		if e.flagVar&FlagMapping[key] == FlagMapping[key] {
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
			tmp.val = util.InputSecretString(tmp.annotation)
		} else {
			tmp.val = util.InputString(tmp.annotation)
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

	if c.Password != "" {
		cipertext, err := config.RowCryptor.Encrypt(c.Password)
		if err != nil {
			log.Fatal(err)
		}
		c.Password = string(cipertext)
	}

	return c
}
