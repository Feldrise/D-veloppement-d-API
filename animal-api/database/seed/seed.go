package seed

import "gorm.io/gorm"

type Seed struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}
