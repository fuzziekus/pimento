package main

import (
	"fmt"

	"github.com/fuzziekus/pimento/hoge/db"
)

func main() {
	printAllCredential()
	for _, d := range [...]string{"aaa1", "aaa2", "aaa3"} {
		createCredential(d)
	}
	printAllCredential()
}

func printAllCredential() {
	credentials := db.NewCredentialRepository().GetAll()
	for _, c := range credentials {
		fmt.Printf("%s\t%s\t%s\n", c.Description, c.UserID, c.Memo)
	}
}

func createCredential(domain string) {
	db.NewCredentialRepository().Create(domain, "bbb", "ccc", "dddd")
}
