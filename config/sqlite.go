package config

import "gorm.io/gorm"

func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	gorm.Open()
}
