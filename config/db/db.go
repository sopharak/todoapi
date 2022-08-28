package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb(conn string) (*gorm.DB, error) {
	connDB, err := gorm.Open(sqlite.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %v", err)
	}
	return connDB, nil
}
