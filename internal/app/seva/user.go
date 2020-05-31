package seva

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrPasswordsDoNotMatch   = errors.New("passwords do not match")
	ErrMissingFirstName      = errors.New("firstName cannot be blank")
	ErrMissingLastName       = errors.New("lastName cannot be blank")
	ErrMissingEmail          = errors.New("email cannot be blank")
	ErrMissingPassword       = errors.New("password cannot be blank")
	ErrMissingPasswordVerify = errors.New("passwordVerify cannot be blank")
	ErrPasswordLength        = errors.New("password must be atleast 8 characters in length")
)

// User will contain information about a user
type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	JWT       string `json:"jwt,omitempty" bson:"-"`
}

// NewUser contains information about a new user
type NewUser struct {
	FirstName      string    `json:"firstName" bson:"firstName"`
	LastName       string    `json:"lastName" bson:"lastName"`
	Email          string    `json:"email" bson:"email"`
	Password       string    `json:"password" bson:"password"`
	PasswordVerify string    `json:"passwordVerify" bson:"-"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
}

// Validate will ensure that the users firstName, lastName, email, and password meet the minimum criteria.
func (nu *NewUser) Validate() error {
	if nu.FirstName == "" {
		return ErrMissingFirstName
	}
	if nu.LastName == "" {
		return ErrMissingLastName
	}
	if nu.Email == "" {
		return ErrMissingEmail
	}
	if nu.Password == "" {
		return ErrMissingPassword
	}
	if nu.PasswordVerify == "" {
		return ErrMissingPasswordVerify
	}
	if nu.Password != nu.PasswordVerify {
		return ErrPasswordsDoNotMatch
	}
	err := validatePassword(nu.Password)
	if err != nil {
		return fmt.Errorf("password failed to meet minimum criteria: %w", err)
	}
	return nil
}

func validatePassword(pass string) error {
	if len(pass) < 8 {
		return ErrPasswordLength
	}
	//contain at least one symbol
	return nil
}
