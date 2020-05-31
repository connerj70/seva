package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/connerj70/seva/cmd/seva/internal/handlers"
	"github.com/connerj70/seva/internal/app/seva"
	"github.com/julienschmidt/httprouter"
)

func TestRetrieveUserSuccess(t *testing.T) {
	// Arrange
	jwtSecretKey := "aa"
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	res := httptest.NewRecorder()
	db := client.Database("test")
	user := handlers.User{JWTSecretKey: jwtSecretKey, DB: db}
	param := httprouter.Param{
		Key:   "id",
		Value: "5ecab985288a64faaa7742fc",
	}
	expectedUser := seva.User{
		FirstName: "conner",
		LastName:  "jensen",
		Email:     "conner@gmail.com",
	}

	// Act
	user.Retrieve(res, req, httprouter.Params{param})
	// Assert
	if res.Code != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		t.Fatalf("wanted a status code of %d, but got %d, with a message of %s", 200, res.Code, body)
	}
	var u seva.User
	err := json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		t.Fatalf("failed to unmarshal response")
	}

	if u.FirstName != expectedUser.FirstName {
		t.Errorf("wanted a first name of %s, but got %s", expectedUser.FirstName, u.FirstName)
	}
	if u.LastName != expectedUser.LastName {
		t.Errorf("wanted a last name of %s, but got %s", expectedUser.LastName, u.LastName)
	}
	if u.Email != expectedUser.Email {
		t.Errorf("wanted an email of %s, but got %s", expectedUser.Email, u.Email)
	}
}

func TestCreateUserSuccess(t *testing.T) {
	// Arrange
	jwtSecretKey := "aa"
	now := time.Now().Unix()
	body := fmt.Sprintf(`{
		"firstName": "conner",
		"lastName": "jensen",
		"email": "test%d@gmail.com",
		"password": "coolkat79",
		"passwordVerify": "coolkat79"
	}`, now)
	bodyReader := bytes.NewReader([]byte(body))
	req := httptest.NewRequest(http.MethodPost, "/user", bodyReader)
	res := httptest.NewRecorder()
	db := client.Database("test")
	user := handlers.User{JWTSecretKey: jwtSecretKey, DB: db}

	// Act
	user.Create(res, req, httprouter.Params{})
	// Assert
	if res.Code != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		t.Fatalf("wanted a status code of %d, but got %d, with a message of %s", 200, res.Code, body)
	}
	var returnValue struct {
		ID string `json:"id"`
	}
	err := json.NewDecoder(res.Body).Decode(&returnValue)
	if err != nil {
		t.Fatalf("failed to unmarshal response")
	}

	if returnValue.ID == "" {
		t.Errorf("wanted a non empty id, but got %q", returnValue.ID)
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	// Arrange
	jwtSecretKey := "aa"
	body := `{
		"id": "5ecab985288a64faaa7742fc",
		"firstName": "conner_updated",
		"lastName": "jensen_updated"
	}`
	bodyReader := bytes.NewReader([]byte(body))
	req := httptest.NewRequest(http.MethodPut, "/user", bodyReader)
	res := httptest.NewRecorder()
	db := client.Database("test")
	user := handlers.User{JWTSecretKey: jwtSecretKey, DB: db}

	// Act
	user.Update(res, req, httprouter.Params{})
	// Assert
	if res.Code != 204 {
		body, _ := ioutil.ReadAll(res.Body)
		t.Fatalf("wanted a status code of %d, but got %d, with a message of %s", 200, res.Code, body)
	}
}
