package model

import (
	"errors"
	"net/http"
)

type User struct {
	ID        uint   `json:"id"`         // @id
	CreatedAt string `json:"created_at"` // the user's creation date

	Email string `json:"email"` // the user's email

	Profile *UserProfile `json:"profile"` // the user's profile
} // @name User

type RegisterPostPayload struct {
	Email     *string `json:"email" validate:"required" example:"admin@feldrise.com"` // the user's email
	Password  *string `json:"password" validate:"required" example:"password"`        // the user's password
	FirstName *string `json:"first_name" example:"Victor"`                            // the user's first name
	LastName  *string `json:"last_name" example:"DENIS"`                              // the user's last name
} // @name RegisterPostPayload

func (rp *RegisterPostPayload) Bind(r *http.Request) error {
	if rp.Email == nil {
		return errors.New("missing email property")
	}

	if rp.Password == nil {
		return errors.New("missing password property")
	}

	if rp.FirstName == nil {
		return errors.New("missing first_name property")
	}

	if rp.LastName == nil {
		return errors.New("missing last_name property")
	}

	return nil
}

type LoginPostPayload struct {
	Email    *string `json:"email" validate:"required" example:"admin@feldrise.com"` // the user's email
	Password *string `json:"password" validate:"required" example:"password"`        // the user's password
} // @name LoginPostPayload

func (lp *LoginPostPayload) Bind(r *http.Request) error {
	if lp.Email == nil {
		return errors.New("missing email property")
	}

	if lp.Password == nil {
		return errors.New("missing password property")
	}

	return nil
}

type ForgotPasswordPostPayload struct {
	Email *string `json:"email" validate:"required" example:"admin@feldrise.com"` // the user's email
} // @name ForgotPasswordPostPayload

func (fp *ForgotPasswordPostPayload) Bind(r *http.Request) error {
	if fp.Email == nil {
		return errors.New("missing email property")
	}

	return nil
}

type ResetPasswordPostPayload struct {
	Email    *string `json:"email" validate:"required" example:"admin@feldrise.com"` // the user's email
	Password *string `json:"password" validate:"required" example:"MyNewPassword"`   // the user's new password
	Code     *string `json:"code" validate:"required" example:"123456"`              // the verification code
} // @name ResetPasswordPostPayload

func (rp *ResetPasswordPostPayload) Bind(r *http.Request) error {
	if rp.Email == nil {
		return errors.New("missing email property")
	}

	if rp.Password == nil {
		return errors.New("missing password property")
	}

	if rp.Code == nil {
		return errors.New("missing code property")
	}

	return nil
}
