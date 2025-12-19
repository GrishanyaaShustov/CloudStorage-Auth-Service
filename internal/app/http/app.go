package http

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	log *slog.Logger
	srv *http.Server
}

type Deps struct {
	Log     *slog.Logger
	Addr    string
	Handler http.Handler

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func New(deps Deps) *App {
	readTO := deps.ReadTimeout
	if readTO == 0 {
		readTO = 5 * time.Second
	}
	writeTO := deps.WriteTimeout
	if writeTO == 0 {
		writeTO = 10 * time.Second
	}
	idleTO := deps.IdleTimeout
	if idleTO == 0 {
		idleTO = 60 * time.Second
	}

	return &App{
		log: deps.Log,
		srv: &http.Server{
			Addr:         deps.Addr,
			Handler:      deps.Handler,
			ReadTimeout:  readTO,
			WriteTimeout: writeTO,
			IdleTimeout:  idleTO,
		},
	}
}

func (a *App) Run() error {
	a.log.Info("http server started", "addr", a.srv.Addr)
	err := a.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (a *App) Shutdown(ctx context.Context) error {
	a.log.Info("http server shutdown")
	return a.srv.Shutdown(ctx)
}
