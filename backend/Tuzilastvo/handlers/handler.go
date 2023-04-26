package handlers

import (
	"Tuzilastvo/data"
	"Tuzilastvo/db"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyPrijava struct{}

type KeyTuzilastvo struct{}

type KeyOptuznica struct{}

const unableToConvertToJson = "Unable to convert to json"

const unableToFindPrijava = "Unable to find prijava."

func NewHandler(l *log.Logger, ur db.Repo) *Handler {
	return &Handler{l, ur}
}

func (u *Handler) CreateTuzilastvo(rw http.ResponseWriter, h *http.Request) {
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

func (u *Handler) GetTuzilastva(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetTuzilastva()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *Handler) GetTuzilastvo(rw http.ResponseWriter, h *http.Request) {
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

func (u *Handler) CreatePrijava(rw http.ResponseWriter, h *http.Request) {
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

func (u *Handler) GetPrijave(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetPrijave(false)

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *Handler) GetJavnePrijave(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetPrijave(true)

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *Handler) GetPrijava(rw http.ResponseWriter, h *http.Request) {
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

func (u *Handler) ConfirmPrijava(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t := u.repo.ConfirmPrijava(id)

	if t == false {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *Handler) DeclinePrijava(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var id = vars["id"]
	t := u.repo.DeclinePrijava(id)

	if t == false {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *Handler) GetOptuznice(rw http.ResponseWriter, h *http.Request) {
	user := u.repo.GetOptuznice()

	err := user.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *Handler) CreateOptuznica(rw http.ResponseWriter, h *http.Request) {
	optuznica := h.Context().Value(KeyOptuznica{}).(*data.Optuznica)

	_, err := u.repo.GetOptuznica(optuznica.KrivicnaPrijava.Id)
	if err == nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if u.repo.CreateOptuznica(optuznica) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *Handler) MiddlewareOptuznicaValidation(next http.Handler) http.Handler {
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

func (u *Handler) MiddlewareTuzilastvoValidation(next http.Handler) http.Handler {
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
