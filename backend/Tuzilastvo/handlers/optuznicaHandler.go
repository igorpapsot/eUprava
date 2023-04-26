package handlers

import (
	"Tuzilastvo/data"
	"Tuzilastvo/db"
	"context"
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
	user := u.repo.GetOptuznice()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *OptuznicaHandler) CreateOptuznica(rw http.ResponseWriter, h *http.Request) {
	optuznica := h.Context().Value(KeyOptuznica{}).(*data.Optuznica)

	prijava, err := u.repo.GetPrijava(optuznica.KrivicnaPrijava.Id)
	optuznica.KrivicnaPrijava = prijava
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	t := u.repo.ConfirmPrijava(optuznica.KrivicnaPrijava.Id)
	if t == false {
		rw.WriteHeader(http.StatusConflict)
		return
	}

	if u.repo.CreateOptuznica(optuznica) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
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
