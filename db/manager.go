package db

import (
	"log"

	"github.com/fuzziekus/pimento/config"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type manager struct {
	db *gorm.DB
}

var isInitialized bool
var mgr manager

func newManager() manager {
	db, err := gorm.Open("sqlite3", config.Mgr().Db.Path+".temp")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	db.AutoMigrate(&Credential{})
	return manager{db: db}
}

func Mgr() manager {
	if !isInitialized {
		mgr = newManager()
		isInitialized = true
	}
	return mgr
}

func (m manager) Close() {
	m.db.Close()
}
