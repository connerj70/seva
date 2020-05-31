package seva_test

import (
	"errors"
	"testing"

	"github.com/connerj70/seva/internal/app/seva"
)

func TestNewUserValidateReturnsCorrectErrorWhenMissingFirstName(t *testing.T) {
	// Arrange
	nu := seva.NewUser{
		FirstName:      "",
		LastName:       "jensen",
		Email:          "conner@gmail.com",
		Password:       "aaa",
		PasswordVerify: "aaa",
	}
	// Act
	err := nu.Validate()
	// Assert
	if !errors.Is(err, seva.ErrMissingFirstName) {
		t.Errorf("wanted an error of %s, but got %s", seva.ErrMissingFirstName, err)
	}
}

func TestNewUserValidateReturnsCorrectErrorWhenMissingLastName(t *testing.T) {
	// Arrange
	nu := seva.NewUser{
		FirstName:      "conner",
		LastName:       "",
		Email:          "conner@gmail.com",
		Password:       "aaa",
		PasswordVerify: "aaa",
	}
	// Act
	err := nu.Validate()
	// Assert
	if !errors.Is(err, seva.ErrMissingLastName) {
		t.Errorf("wanted an error of %s, but got %s", seva.ErrMissingLastName, err)
	}
}

func TestNewUserValidateReturnsCorrectErrorWhenMissingEmail(t *testing.T) {
	// Arrange
	nu := seva.NewUser{
		FirstName:      "conner",
		LastName:       "jensen",
		Email:          "",
		Password:       "aaa",
		PasswordVerify: "aaa",
	}
	// Act
	err := nu.Validate()
	// Assert
	if !errors.Is(err, seva.ErrMissingEmail) {
		t.Errorf("wanted an error of %s, but got %s", seva.ErrMissingEmail, err)
	}
}

func TestNewUserValidateReturnsCorrectErrorWhenMissingPassword(t *testing.T) {
	// Arrange
	nu := seva.NewUser{
		FirstName:      "conner",
		LastName:       "jensen",
		Email:          "conner@gmail.com",
		Password:       "",
		PasswordVerify: "aaa",
	}
	// Act
	err := nu.Validate()
	// Assert
	if !errors.Is(err, seva.ErrMissingPassword) {
		t.Errorf("wanted an error of %s, but got %s", seva.ErrMissingPassword, err)
	}
}

func TestNewUserValidateReturnsCorrectErrorWhenMissingPasswordVerify(t *testing.T) {
	// Arrange
	nu := seva.NewUser{
		FirstName:      "conner",
		LastName:       "jensen",
		Email:          "conner@gmail.com",
		Password:       "aaa",
		PasswordVerify: "",
	}
	// Act
	err := nu.Validate()
	// Assert
	if !errors.Is(err, seva.ErrMissingPasswordVerify) {
		t.Errorf("wanted an error of %s, but got %s", seva.ErrMissingPasswordVerify, err)
	}
}

func TestValidatePasswordReturnsCorrectErrorWhenPasswordIsLessThanEightCharacters(t *testing.T) {
	// Arrange
	nu := seva.NewUser{
		FirstName:      "conner",
		LastName:       "jensen",
		Email:          "conner@gmail.com",
		Password:       "aaa",
		PasswordVerify: "aaa",
	}
	// Act
	err := nu.Validate()
	// Assert
	if !errors.Is(err, seva.ErrPasswordLength) {
		t.Errorf("wanted an error of %s, but got %s", seva.ErrPasswordLength, err)
	}
}
