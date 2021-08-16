package db

import (
	// "crypto"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fuzziekus/pimento/config"
)

type Credential struct {
	ID        int
	ItemName  string `gorm:"unique;not null" csv:"item_name"`
	UserName  string `csv:"user_name"`
	Password  string `csv:"password"`
	Tag       string `csv:"tag"`
	CreatedAt time.Time
	UpdateAt  sql.NullTime
}

type Credentials []Credential

type credentialRepository struct{}

func NewCredentialRepository() credentialRepository {
	return credentialRepository{}
}

func (r credentialRepository) CreateWithRawVal(item_name, user_name, password, tag string) {
	cipertext, err := config.RowCryptor.Encrypt(password)
	if err != nil {
		log.Fatal(err)
	}
	credential := Credential{
		ItemName: item_name,
		UserName: user_name,
		Password: string(cipertext),
		Tag:      tag,
	}
	res := Mgr().db.Create(&credential)
	if res.Error != nil {
		fmt.Fprintln(os.Stderr, res.Error)
	}
}

func (r credentialRepository) Create(c Credential) {
	res := Mgr().db.Create(&c)
	if res.Error != nil {
		fmt.Fprintln(os.Stderr, res.Error)
	}
}

func (r credentialRepository) GetAll() Credentials {
	credentials := Credentials{}
	Mgr().db.Find(&credentials)
	return credentials
}

func (r credentialRepository) GetSingleRowByItemName(item_name string) Credential {
	credential := Credential{
		// ItemName: item_name,
	}
	if err := Mgr().db.First(&credential, "item_name = ?", item_name).Error; err != nil {
		log.Fatalf("対象のレコードが見つかりませんでした")
	}

	return credential
}

func (r credentialRepository) GetSingleRowById(id int) Credential {
	credential := Credential{
		ID: id,
	}
	Mgr().db.First(&credential)
	return credential
}

func (r credentialRepository) UpdateRow(id int, c Credential) Credential {
	credential := Credential{
		ID: id,
	}
	Mgr().db.Model(&credential).Update(&c)
	return credential
}

func (r credentialRepository) DeleteRow(id int) Credential {
	credential := Credential{
		ID: id,
	}
	Mgr().db.Delete(&credential)
	return credential
}
