package model

type UserProfile struct {
	FirstName            *string `json:"first_name"`
	LastName             *string `json:"last_name"`
	SocialSecurityNumber *string `json:"social_security_number"`
} // @name UserProfile
