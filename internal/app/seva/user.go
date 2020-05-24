package seva

import (
	"fmt"
	"time"
)

// User will contain information about a user
type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password,omitempty" bson:"password"`
	JWT       string `json:"jwt,omitempty"`
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

func (nu *NewUser) Validate() error {
	if nu.FirstName == "" {
		return fmt.Errorf("firstName cannot be blank")
	}
	if nu.LastName == "" {
		return fmt.Errorf("lastName cannot be blank")
	}
	if nu.Email == "" {
		return fmt.Errorf("email cannot be blank")
	}
	if nu.Password == "" {
		return fmt.Errorf("password cannot be blank")
	}
	if nu.PasswordVerify == "" {
		return fmt.Errorf("passwordVerify cannot be blank")
	}
	if nu.Password != nu.PasswordVerify {
		return fmt.Errorf("passwords do not match")
	}
	err := validatePassword(nu.Password)
	if err != nil {
		return fmt.Errorf("password failed to meet minimum criteria: %w", err)
	}
	return nil
}

func validatePassword(pass string) error {
	if len(pass) < 8 {
		return fmt.Errorf("password must be atleast 8 characters in length")
	}
	return nil
}
