package handlers

import (
	"Tuzilastvo/data"
	"Tuzilastvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PrijavaHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyPrijava struct{}

const unableToConvertToJson = "Unable to convert to json"

func NewPrijavaHandler(l *log.Logger, ur db.Repo) *PrijavaHandler {
	return &PrijavaHandler{l, ur}
}

func (u *PrijavaHandler) CreatePrijava(rw http.ResponseWriter, h *http.Request) {
	prijava := h.Context().Value(KeyPrijava{}).(*data.KrivicnaPrijava)

	_, err := u.repo.GetPrijava(prijava.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreatePrijava(prijava) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *PrijavaHandler) GetPrijave(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetPrijave(false)

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrijavaHandler) SearchPrijave(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var input = vars["input"]
	user := u.repo.SearchPrijave(input)

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrijavaHandler) GetJavnePrijave(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetPrijave(true)

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrijavaHandler) GetPrijava(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t, err := u.repo.GetPrijava(id)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrijavaHandler) ConfirmPrijava(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]

	prijava, _ := u.repo.GetPrijava(id)
	if prijava.Status != 2 {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	t := u.repo.ConfirmPrijava(id)
	if t == false {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrijavaHandler) DeclinePrijava(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]

	prijava, _ := u.repo.GetPrijava(id)
	if prijava.Status != 2 {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	t := u.repo.DeclinePrijava(id)
	if t == false {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrijavaHandler) MiddlewarePrijavaValidation(next http.Handler) http.Handler {
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

func (u *PrijavaHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
