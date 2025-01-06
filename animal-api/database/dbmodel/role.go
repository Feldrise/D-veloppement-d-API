package dbmodel

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}
