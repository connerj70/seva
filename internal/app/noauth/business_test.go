package noauth

import (
	"fmt"
	"testing"
)

var b = new(Business)

func TestRegister(t *testing.T) {
	user := &User{Email: "connerj70@gmail.com", Password: "coolcat772"}
	err := b.Register(user)
	fmt.Println(err)
	t.Error("hi")
}
