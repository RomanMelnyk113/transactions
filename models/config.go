package models

import "gorm.io/gorm"

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&Transaction{}, &Account{})
}
