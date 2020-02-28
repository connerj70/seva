package noauth

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type serviceAdapter interface {
	Register(*User) error
	GetUserByEmail(string) (*User, error)
}

// Business will handle all the business logic
type Business struct {
	Service serviceAdapter
}

// Register will handle registering a user
func (b *Business) Register(user *User) error {
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	// Check if there is already a user with this email in our database
	userCheck, err := b.Service.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if userCheck.Email != "" {
		return errors.New("a user with this email already exists")
	}

	if user.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	// Hash the users password
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	passwordSum := hash.Sum(nil)
	passwordHashString := fmt.Sprintf("%x", passwordSum)

	user.Password = passwordHashString

	err = b.Service.Register(user)
	if err != nil {
		return err
	}

	return nil
}

// LogIn will attempt to log the user in
func (b *Business) LogIn(user *User) (err error) {

	verifyUser, err := b.Service.GetUserByEmail(user.Email)
	if err != nil {
		return
	}

	// Hash the users password
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	passwordSum := hash.Sum(nil)
	passwordHashString := fmt.Sprintf("%x", passwordSum)

	// Check to see if the passwords match
	if passwordHashString != verifyUser.Password {
		err = errors.New("Password does not match")
		return
	}

	// Set up and generate jwt token
	sampleKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(sampleKey)
	if err != nil {
		return
	}
	user.JWT = tokenString

	return nil
}
