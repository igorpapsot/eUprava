package handlers

import (
	"Tuzilastvo/data"
	"Tuzilastvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TuzilastvoHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyTuzilastvo struct{}

func NewTuzilastvoHandler(l *log.Logger, ur db.Repo) *TuzilastvoHandler {
	return &TuzilastvoHandler{l, ur}
}

func (u *TuzilastvoHandler) CreateTuzilastvo(rw http.ResponseWriter, h *http.Request) {
	tuzilastvo := h.Context().Value(KeyTuzilastvo{}).(*data.Tuzilastvo)

	_, err := u.repo.GetTuzilastvo(tuzilastvo.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreateTuzilastvo(tuzilastvo) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *TuzilastvoHandler) GetTuzilastva(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetTuzilastva()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *TuzilastvoHandler) GetTuzilastvo(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t, err := u.repo.GetTuzilastvo(id)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *TuzilastvoHandler) MiddlewareTuzilastvoValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		tuzilastvo := &data.Tuzilastvo{}
		err := tuzilastvo.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyTuzilastvo{}, tuzilastvo)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
