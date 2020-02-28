package app

import (
	"github.com/connerj70/seva/internal/app/noauth"
	"github.com/connerj70/seva/internal/app/sandwich"
	"github.com/connerj70/seva/internal/connection"
)

// WireNoAuth sets up the 3 tier architecture for noauth
func WireNoAuth() noauth.Receiver {
	s := &noauth.Service{DB: connection.Mongo}
	b := &noauth.Business{Service: s}
	r := noauth.Receiver{Business: b}
	return r
}

// WireSandwich setups up the 3 tier architecture for sandwich
func WireSandwich() sandwich.Receiver {
	s := &sandwich.Service{DB: connection.Mongo}
	b := &sandwich.Business{Service: s}
	r := &sandwich.Receiver{Business: b}
	return r
}
