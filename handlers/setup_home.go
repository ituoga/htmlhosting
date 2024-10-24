package handlers

import (
	"net/http"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/ituoga/go-start/web/views/homeview"
)

func SetupHome(router chi.Router, session sessions.Store, ns *embeddednats.Server) error {

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		homeview.Index().Render(r.Context(), w)
	}

	router.Get("/", homeHandler)

	return nil
}
