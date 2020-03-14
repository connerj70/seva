package noauth

import (
	"testing"
)

type MockService struct{}

var b = Business{Service: &MockService{}}

func TestRegister(t *testing.T) {
	user := &User{Email: "connerj70@gmail.com", Password: "coolcat772"}
	err := b.Register(user)
	if err != nil {
		t.Errorf("wanted an error of nil, but got %s", err)
	}
}

func (ms *MockService) GetUserByEmail(string) (*User, error) {
	return &User{}, nil
}
func (ms *MockService) Register(*User) error {
	return nil
}
