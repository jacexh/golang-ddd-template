package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jacexh/golang-ddd-template/internal/logger"
	"github.com/jacexh/golang-ddd-template/internal/option"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	"github.com/jacexh/golang-ddd-template/router/api"
)

func BuildRouter(option option.RouterOption) http.Handler {
	r := chi.NewRouter()

	if option.Timeout != 0 {
		infection.DefaultTimeout = time.Duration(option.Timeout) * time.Second
	}
	r.Use(GlobalTimeout)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(ZapLog(logger.Logger))
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	{
		r.Post("/api/v1/user", api.CreateUser)
	}
	return r
}
