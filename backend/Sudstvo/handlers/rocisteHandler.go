package handlers

import (
	"Sudstvo/data"
	"Sudstvo/db"
	"context"
	"log"
	"net/http"
)

type RocisteHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyRociste struct{}

func NewRocisteHandler(l *log.Logger, ur db.Repo) *RocisteHandler {
	return &RocisteHandler{l, ur}
}

func (u *RocisteHandler) GetRocista(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetRocista()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *RocisteHandler) CreateRociste(rw http.ResponseWriter, h *http.Request) {
	rociste := h.Context().Value(KeyRociste{}).(*data.Rociste)

	_, err := u.repo.GetRociste(rociste.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreateRociste(rociste) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *RocisteHandler) MiddlewareRocisteValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		rociste := &data.Rociste{}
		err := rociste.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyRociste{}, rociste)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
