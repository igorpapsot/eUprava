package handlers

import (
	"Sudstvo/data"
	"Sudstvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type OptuznicaHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyOptuznica struct{}

func NewOptuznicaHandler(l *log.Logger, ur db.Repo) *OptuznicaHandler {
	return &OptuznicaHandler{l, ur}
}

func (u *OptuznicaHandler) GetOptuznice(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetPoternice()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *OptuznicaHandler) GetOptuznica(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t, err := u.repo.GetOptuznica(id)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *OptuznicaHandler) MiddlewareOptuznicaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		optuznica := &data.Optuznica{}
		err := optuznica.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyOptuznica{}, optuznica)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
