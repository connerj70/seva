package noauth

// User will hold user data
type User struct {
	Email    string
	Password string `json:",omitempty"`
	JWT      string
}
