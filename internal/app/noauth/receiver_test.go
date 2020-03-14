package noauth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockBusiness struct{}

var rec = Receiver{Business: &MockBusiness{}}

func TestError500WithBadJSON(t *testing.T) {
	postJSON := `{"Email: "connerj70@gmail.com", "Password": "akslnva;s092"}`
	r := httptest.NewRequest(http.MethodPost, "/noauth/register", strings.NewReader(postJSON))
	r.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()

	rec.Register(w, r)

	resp := w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("wanted status code of: %v but got: %v given: %s", 500, resp.StatusCode, postJSON)
	}
}

func TestErrorFromBusiness(t *testing.T) {
	postJSON := `{"Email": "fail@fail.fail", "Password": "akslnva;s092"}`
	r := httptest.NewRequest(http.MethodPost, "/noauth/register", strings.NewReader(postJSON))
	r.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()

	rec.Register(w, r)

	resp := w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("wanted status code of: %v but got: %v given: %s", 500, resp.StatusCode, postJSON)
	}
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != "error from business" {
		t.Errorf("wanted body of: %s but got: %s given: %s", "error from business", body, postJSON)
	}
}

func (mb *MockBusiness) Register(user *User) error {
	if user.Email == "fail@fail.fail" {
		return fmt.Errorf("error from business")
	}
	return nil
}

func (mb *MockBusiness) LogIn(user *User) (err error) {
	return nil
}
