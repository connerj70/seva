package sandwich

import (
	"errors"
	"github.com/connerj70/seva/internal/cerr"
	"testing"
)

func TestBusinessPostReturnsCorrectErrorWhenSandwichValidationFails(t *testing.T) {
	// Arrange
	f := false
	s := Sandwich{Delicious: &f}
	b := Business{}
	// Act
	err := b.Post(&s)
	// Assert
	if err == nil {
		t.Fatalf("wanted a non nil error, but got nil")
	}
	expectedErr := "sandwich must have a delicious property of True"
	cErr := &cerr.InternalError{}
	if errors.As(err, cErr) {
		if cErr.Detail != expectedErr {
			t.Errorf(`wanted an error message of "%s", but got "%s"`, expectedErr, cErr.Detail)
		}
	}
}
