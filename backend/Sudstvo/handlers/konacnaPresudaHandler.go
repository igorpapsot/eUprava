package handlers

import (
	"Sudstvo/data"
	"Sudstvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type KonacnaPresudaHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyKonacnaPresuda struct{}

func NewKonacnaPresudaHandler(l *log.Logger, ur db.Repo) *KonacnaPresudaHandler {
	return &KonacnaPresudaHandler{l, ur}
}

func (u *KonacnaPresudaHandler) CreateKonacnaPresuda(rw http.ResponseWriter, h *http.Request) {
	konacnaPresuda := h.Context().Value(KeyKonacnaPresuda{}).(*data.KonacnaPresuda)

	_, err := u.repo.GetKonacnaPresuda(konacnaPresuda.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreateKonacnaPresuda(konacnaPresuda) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *KonacnaPresudaHandler) GetKonacnaPresuda(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t, err := u.repo.GetKonacnaPresuda(id)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *KonacnaPresudaHandler) MiddlewareKonacnaPresudaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		konacnaPresuda := &data.KonacnaPresuda{}
		err := konacnaPresuda.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyKonacnaPresuda{}, konacnaPresuda)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
