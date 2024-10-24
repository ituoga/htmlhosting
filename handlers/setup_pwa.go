package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupManifest(router chi.Router, session sessions.Store, ns *embeddednats.Server) error {

	type Icons struct {
		Src     string `json:"src"`
		Type    string `json:"type"`
		Sizes   string `json:"sizes"`
		Purpose string `json:"purpose"`
	}

	type Params struct {
		Title string `json:"title"`
		Text  string `json:"text"`
		URL   string `json:"url"`
	}

	type ShareTarget struct {
		Action string `json:"action"`
		Method string `json:"method"`
		Params Params `json:"params"`
	}

	type Manifest struct {
		Version         string      `json:"version"`
		ShortName       string      `json:"short_name"`
		Name            string      `json:"name"`
		ShareTarget     ShareTarget `json:"share_target"`
		StartURL        string      `json:"start_url"`
		BackgroundColor string      `json:"background_color"`
		Display         string      `json:"display"`
		ThemeColor      string      `json:"theme_color"`
		Icons           []Icons     `json:"icons"`
	}

	router.Get("/manifest.json", func(w http.ResponseWriter, r *http.Request) {

		var manifest Manifest

		manifest.Version = "1.3"
		manifest.ShortName = "App Dev"
		manifest.Name = "App Dev"
		manifest.ShareTarget.Action = "/phone-share"
		manifest.ShareTarget.Method = "GET"
		manifest.ShareTarget.Params.Title = "title"
		manifest.ShareTarget.Params.Text = "text"
		manifest.ShareTarget.Params.URL = "url"
		manifest.StartURL = "/"
		manifest.BackgroundColor = "#ffffff"
		manifest.Display = "standalone"
		manifest.ThemeColor = "#0d085c"
		manifest.Icons = []Icons{
			{
				Src:     "/static/favicon.png",
				Type:    "image/png",
				Sizes:   "512x512",
				Purpose: "any maskable",
			},
		}

		b, err := json.Marshal(manifest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	})
	return nil
}
