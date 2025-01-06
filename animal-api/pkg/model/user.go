package model

import (
	"errors"
	"net/http"
)

type UserCreatePayload struct {
	Email     *string `json:"email" validate:"required,email" example:"admin@feldrise.com"`
	FirstName *string `json:"first_name" validate:"required" example:"Victor"`
	LastName  *string `json:"last_name" validate:"required" example:"DENIS"`
} // @name InternalUserCreatePayload

func (iup *UserCreatePayload) Bind(r *http.Request) error {
	if iup.Email == nil {
		return errors.New("missing email property")
	}
	if iup.FirstName == nil {
		return errors.New("missing first_name property")
	}
	if iup.LastName == nil {
		return errors.New("missing last_name property")
	}

	return nil
}
