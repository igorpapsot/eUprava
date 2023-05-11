package handlers

import (
	"Sudstvo/data"
	"Sudstvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PoternicaHandler struct {
	logger *log.Logger
	repo   db.Repo
}

const unableToConvertToJson = "Unable to convert to json"

type KeyPoternica struct{}

func NewPoternicaHandler(l *log.Logger, ur db.Repo) *PoternicaHandler {
	return &PoternicaHandler{l, ur}
}

func (u *PoternicaHandler) GetPoternice(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetPoternice()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PoternicaHandler) GetPoternica(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t, err := u.repo.GetPoternica(id)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PoternicaHandler) CreatePoternica(rw http.ResponseWriter, h *http.Request) {
	poternica := h.Context().Value(KeyPoternica{}).(*data.Poternica)

	_, err := u.repo.GetPoternica(poternica.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreatePoternica(poternica) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *PoternicaHandler) MiddlewarePoternicaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		poternica := &data.Poternica{}
		err := poternica.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyPoternica{}, poternica)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
