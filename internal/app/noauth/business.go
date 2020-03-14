package noauth

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/connerj70/seva/internal/cerr"
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
		ie := &cerr.InternalError{
			Header: "user verification",
			Detail: "a user with this email already exists",
		}
		ie.SetStatusCode(http.StatusConflict)
		return ie
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
		var ie *cerr.InternalError
		if errors.As(err, &ie) {
			return &cerr.InternalError{
				Header: fmt.Sprintf("register business %s", ie.Header),
				Detail: fmt.Sprintf("there was a problem registering the user %s", ie.Detail),
				Err:    ie.Err,
			}
		}
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

	// Remove the password from the returning user struct
	user.Password = ""

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
