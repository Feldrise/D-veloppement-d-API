package seed

import (
	"feldrise.com/animal-api/database/dbmodel"
	"gorm.io/gorm"
)

func SeedV1(database *gorm.DB) error {
	// Roles
	roleAdmin := dbmodel.Role{
		Model: gorm.Model{
			ID: 1,
		},
		Name:        "admin",
		Description: "Administrator",
	}

	roleEditor := dbmodel.Role{
		Model: gorm.Model{
			ID: 2,
		},
		Name:        "veterinaire",
		Description: "Veterinaire",
	}

	roleLabel := dbmodel.Role{
		Model: gorm.Model{
			ID: 3,
		},
		Name:        "client",
		Description: "Client",
	}

	database.Create(&roleAdmin)
	database.Create(&roleEditor)
	database.Create(&roleLabel)

	return nil
}
