package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/transinterdigital/internal/config"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/internal/gateway/handlers/middleware"
)

func NewHandlers(usecases usecases.UseCase, config config.Config) *Handlers {
	return &Handlers{
		usecases,
		config,
	}
}

type Handlers struct {
	UseCase usecases.UseCase
	Config  config.Config
}

func (h Handlers) Handlers(r *chi.Mux) *chi.Mux {
	r.Get("/accounts", h.GetListAccount)
	r.Get("/accounts/{account_id}/balance", h.GetBalanceAccount)
	r.Post("/accounts", h.CreateAccount)
	r.Post("/login", h.LoginGetToken)
	r.Get("/transfers", middleware.Authenticate(h.GetListTransfers, h.Config))
	r.Post("/transfers", middleware.Authenticate(h.TransferAccountToAnother, h.Config))
	return r
}
