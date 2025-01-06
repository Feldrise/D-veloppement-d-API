package dbmodel

import (
	"feldrise.com/animal-api/pkg/model"
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	FirstName *string
	LastName  *string

	SocialSecurityNumber *string

	UserID *uint
}

func (userProfile *UserProfile) ToModel() *model.UserProfile {
	return &model.UserProfile{
		FirstName:            userProfile.FirstName,
		LastName:             userProfile.LastName,
		SocialSecurityNumber: userProfile.SocialSecurityNumber,
	}
}
