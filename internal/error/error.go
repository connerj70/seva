package error

// InternalError will store information about errors that happen in our application
type InternalError struct {
	Header string
	Detail string
}
