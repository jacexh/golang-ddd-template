package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jacexh/golang-ddd-template/internal/option"
	"github.com/jacexh/golang-ddd-template/internal/transport/rest/api"
	"github.com/jacexh/golang-ddd-template/pkg/infection"
	chimiddleware "github.com/jacexh/gopkg/chi-middleware"
	"go.uber.org/zap"
)

func BuildRouter(option option.RouterOption, log *zap.Logger) http.Handler {
	r := chi.NewRouter()

	if option.Timeout != 0 {
		infection.DefaultTimeout = time.Duration(option.Timeout) * time.Second
	}
	r.Use(InfectContext)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(chimiddleware.RequestZapLog(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	{
		r.Post("/api/v1/user", api.CreateUser)
	}
	return r
}
