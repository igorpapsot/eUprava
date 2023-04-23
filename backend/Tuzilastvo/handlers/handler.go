package handlers

import (
	"Tuzilastvo/data"
	"Tuzilastvo/db"
	"context"
	"log"
	"net/http"
)

type Handler struct {
	logger   *log.Logger
	userRepo db.Repo
}

type KeyPrijava struct{}

func NewHandler(l *log.Logger, ur db.Repo) *Handler {
	return &Handler{l, ur}
}

func (u *Handler) CreatePrijava(rw http.ResponseWriter, h *http.Request) {

}

func (u *Handler) GetPrijave(rw http.ResponseWriter, h *http.Request) {

}

func (u *Handler) ConfirmPrijava(rw http.ResponseWriter, h *http.Request) {

}

func (u *Handler) DeclinePrijava(rw http.ResponseWriter, h *http.Request) {

}

func (u *Handler) MiddlewarePrijavaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		prijava := &data.KrivicnaPrijava{}
		err := prijava.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyPrijava{}, prijava)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *Handler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
