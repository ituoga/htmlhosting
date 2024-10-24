package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/invopop/ctxi18n"
	"github.com/ituoga/go-start/locales"
)

func SetupRoutes(logger *slog.Logger, router chi.Router) error {
	ns, err := embeddednats.New(context.Background())
	if err != nil {
		return fmt.Errorf("error creating embedded nats server: %w", err)
	}
	ns.WaitForServer()

	if err := ctxi18n.LoadWithDefault(locales.Content, "en"); err != nil {
		panic(err)
	}

	sessionStore := sessions.NewCookieStore([]byte("secret")) // @todo: move to env
	sessionStore.MaxAge(int(24 * time.Hour / time.Second))

	if err := errors.Join(
		SetupManifest(router, sessionStore, ns),
		SetupHome(router, sessionStore, ns),
		// SetupApi(logger, router, sessionStore, ns),
	); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	return nil
}
