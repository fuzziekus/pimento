package ui

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fuzziekus/pimento/config"
	"github.com/fuzziekus/pimento/db"
)

func (e Flags) DisplayHeader() {
	e.calcCondition()
	header := []interface{}{}
	fmtTemplate := []string{}

	for _, key := range ColumnOrder {
		if e.flagVar&FlagMapping[key] == FlagMapping[key] {
			header = append(header, key)
			fmtTemplate = append(fmtTemplate, "%s")
		}
	}

	format := strings.Join(fmtTemplate, "\t")
	if !e.NoHeader {
		fmt.Printf(format+"\n", header...)
	}

}

func (e Flags) DisplaySpecifyColumn(c db.Credential) {
	e.calcCondition()
	// fmt.Println(e.flagVar)
	// 出力する変数を e.flagVar の値から決定する
	targets := []interface{}{}
	fmtTemplate := []string{}

	rtCredential := reflect.TypeOf(c)
	uv := reflect.ValueOf(&c).Elem()
	for i := 0; i < rtCredential.NumField(); i++ {
		f := rtCredential.Field(i)
		for _, key := range ColumnOrder {
			v := uv.FieldByName(key).Interface()
			if f.Name == key && e.flagVar&FlagMapping[key] == FlagMapping[key] {
				fmtTemplate = append(fmtTemplate, "%s")

				if f.Name == "Password" {
					plaintext, err := config.RowCryptor.Decrypt(v.(string))
					if err != nil {
						log.Fatal(err)
					}
					targets = append(targets, string(plaintext))
				} else {
					targets = append(targets, v.(string))
				}
			}
		}
	}

	format := strings.Join(fmtTemplate, "\t")
	fmt.Printf(format+"\n", targets...)
}
