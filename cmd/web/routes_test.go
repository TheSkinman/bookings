package main

import (
	"fmt"
	"testing"

	"gethub.com/TheSkinman/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing: test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}

}
