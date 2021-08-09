package ui

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/fuzziekus/pimento/config"
	"github.com/fuzziekus/pimento/db"
	"github.com/olekukonko/tablewriter"
)

func (e Flags) generateHeadersWithFormat() ([]string, []tablewriter.Colors) {
	headers := []string{}
	formats := []tablewriter.Colors{}

	for _, key := range ColumnOrder {
		if e.flagVar&FlagMapping[key] == FlagMapping[key] {
			headers = append(headers, key)
			formats = append(formats, tablewriter.Colors{tablewriter.Bold})
		}
	}

	return headers, formats
}

// 出力する変数を e.flagVar の値から決定する
func (e Flags) generateRowWithSpecifyColumn(c db.Credential) []string {
	targets := []string{}
	rtCredential := reflect.TypeOf(c)
	uv := reflect.ValueOf(&c).Elem()

	// fmt.Printf("%+v", e)
	for i := 0; i < rtCredential.NumField(); i++ {
		f := rtCredential.Field(i)
		for _, key := range ColumnOrder {
			v := uv.FieldByName(key).Interface()
			if f.Name == key && e.flagVar&FlagMapping[key] == FlagMapping[key] {
				if key == "Password" && !e.NoPass {
					plaintext, err := config.RowCryptor.Decrypt(v.(string))
					if err != nil {
						log.Fatal(err)
					}
					targets = append(targets, string(plaintext))
				} else if key == "Password" && e.NoPass {
					targets = append(targets, "*********")
				} else {
					targets = append(targets, v.(string))
				}
			}
		}
	}

	return targets
}

func (e Flags) DisplayRows(credentials db.Credentials) {
	e.calcCondition()
	data := [][]string{}
	var headers []string
	var formats []tablewriter.Colors

	headers, formats = e.generateHeadersWithFormat()

	for _, c := range credentials {
		data = append(data, e.generateRowWithSpecifyColumn(c))
	}

	if e.FormatTable {
		table := newTableWriter()
		if !e.NoHeader {
			table.SetHeader(headers)
			table.SetHeaderColor(formats...)
		}
		table.AppendBulk(data)
		table.Render()
	} else if e.FormatCSV {
		if !e.NoHeader {
			fmt.Fprintln(os.Stdout, strings.Join(appendDoubleQuote(headers), ","))
		}
		for _, d := range data {
			fmt.Fprintln(os.Stdout, strings.Join(appendDoubleQuote(d), ","))
		}

	}
}

func (e Flags) DisplayRow(c db.Credential) {
	table := newTableWriter()

	if !e.NoHeader {
		headers, formats := e.generateHeadersWithFormat()
		table.SetHeader(headers)
		table.SetHeaderColor(formats...)
	}

	table.Append(e.generateRowWithSpecifyColumn(c))
	table.Render()
}

func newTableWriter() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	return table
}

func appendDoubleQuote(array []string) []string {
	var newslice []string
	for _, a := range array {
		newslice = append(newslice, "\""+a+"\"")
	}
	return newslice
}
