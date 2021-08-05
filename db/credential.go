package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type Credential struct {
	ID          int
	Description string `gorm:"unique;not null"`
	UserId      string
	Password    string
	Memo        string
	CreatedAt   time.Time
	UpdateAt    sql.NullTime
}

type Credentials []Credential

type CredentialRepository struct{}

func NewCredentialRepository() CredentialRepository {
	return CredentialRepository{}
}

func (r CredentialRepository) CreateWithRawVal(description, user_id, password, memo string) {
	credential := Credential{
		Description: description,
		UserId:      user_id,
		Password:    password,
		Memo:        memo,
	}
	res := Mgr().db.Create(&credential)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
}

func (r CredentialRepository) Create(c Credential) {
	res := Mgr().db.Create(&c)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
}

func (r CredentialRepository) GetAll() Credentials {
	credentials := Credentials{}
	Mgr().db.Find(&credentials)
	return credentials
}

func (r CredentialRepository) GetSingleRowByDescription(description string) Credential {
	credential := Credential{
		// Description: description,
	}
	if err := Mgr().db.First(&credential, "description = ?", description).Error; err != nil {
		fmt.Println("対象のレコードが見つかりませんでした")
		os.Exit(1)
	}
	return credential
}

func (r CredentialRepository) GetSingleRowById(id int) Credential {
	credential := Credential{
		ID: id,
	}
	Mgr().db.First(&credential)
	return credential
}

func (r CredentialRepository) UpdateRow(id int, c Credential) Credential {
	credential := Credential{
		ID: id,
	}
	Mgr().db.Model(&credential).Update(&c)
	return credential
}
