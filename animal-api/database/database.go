package database

import (
	"log"

	"feldrise.com/animal-api/database/dbmodel"
	"feldrise.com/animal-api/database/seed"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) {
	database.AutoMigrate(
		&seed.Seed{},
		&dbmodel.User{},
		&dbmodel.Role{},
		&dbmodel.UserProfile{},
	)

	log.Println("Database migrated successfully")
}

func ApplySeeds(database *gorm.DB) {
	seedsToApply := []struct {
		Name     string
		SeedFunc func(*gorm.DB) error
	}{
		{"SeedV1", seed.SeedV1},
	}

	for _, seedToApply := range seedsToApply {
		if !isSeedApplied(database, seedToApply.Name) {
			log.Printf("Applying seed %s", seedToApply.Name)
			if err := seedToApply.SeedFunc(database); err != nil {
				log.Fatalf("Error applying seed %s: %s", seedToApply.Name, err)
			}
			markSeedAsApplied(database, seedToApply.Name)
		}
	}

	log.Println("Seeds applied")
}

func isSeedApplied(database *gorm.DB, name string) bool {
	var count int64
	database.Model(&seed.Seed{}).Where("name = ?", name).Count(&count)

	return count > 0
}

func markSeedAsApplied(database *gorm.DB, name string) {
	database.Create(&seed.Seed{Name: name})
}
