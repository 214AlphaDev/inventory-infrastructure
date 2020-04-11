package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var RandomTestDB = func() (*gorm.DB, error) {

	pathBytes := make([]byte, 30)
	_, err := rand.Read(pathBytes)
	if err != nil {
		panic(err)
	}

	return gorm.Open("sqlite3", "/tmp/"+hex.EncodeToString(pathBytes))

}

var StaticTestDB = func() (*gorm.DB, error) {

	return gorm.Open("sqlite3", "/tmp/static_sqlite3")

}
