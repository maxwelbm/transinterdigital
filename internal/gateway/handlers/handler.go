package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/internal/gateway/api/middleware"
)

func NewHandler(usecases usecases.UseCase) *Handler {
	return &Handler{
		usecases,
	}
}

type Handler struct {
	UseCase usecases.UseCase
}

func (h Handler) Handlers(r *chi.Mux) *chi.Mux {
	r.Get("/accounts", h.GetListAccount)
	r.Get("/accounts/{account_id}/balance", h.GetBalanceAccount)
	r.Post("/accounts", h.CreateAccount)
	r.Post("/login", h.LoginGetToken)
	r.Get("/transfers", middleware.Authenticate(h.GetListTransfers))
	r.Post("/transfers", middleware.Authenticate(h.TransferAccountToAnother))
	return r
}
