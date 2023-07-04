package handlers

import (
	"Sudstvo/data"
	"Sudstvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SudHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeySud struct{}

func NewSudHandler(l *log.Logger, ur db.Repo) *SudHandler {
	return &SudHandler{l, ur}
}

func (u *SudHandler) CreateSud(rw http.ResponseWriter, h *http.Request) {
	sud := h.Context().Value(KeySud{}).(*data.Sud)

	_, err := u.repo.GetSud(sud.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreateSud(sud) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *SudHandler) GetSudovi(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetSudovi()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *SudHandler) GetSud(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t, err := u.repo.GetSud(id)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *SudHandler) MiddlewareSudValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		sud := &data.Sud{}
		err := sud.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeySud{}, sud)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
