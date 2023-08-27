package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type PrelazakHandler struct {
	logger *log.Logger
	repo   db.GpRepo
}

type PrelazakInfoStruct struct {
	PolicajacId string //`json:"policajacId"`
	ProveraId   string //`json:"proveraId"`
}

type PrelazakKey struct{}

func NewPrelazakHandler(l *log.Logger, ur db.GpRepo) *PrelazakHandler {
	return &PrelazakHandler{l, ur}
}

func (p *PrelazakHandler) CreatePrelazakHandler(rw http.ResponseWriter, h *http.Request) {
	decoder := json.NewDecoder(h.Body)
	var prelazakInfo PrelazakInfoStruct
	err := decoder.Decode(&prelazakInfo)

	policajac, err := p.repo.GetPolicajac(prelazakInfo.PolicajacId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to find policajac.", err)
		return
	}
	provera, err := p.repo.GetProveraById(prelazakInfo.ProveraId)
	if err != nil {
		http.Error(rw, "No provera found", http.StatusNotAcceptable)
		p.logger.Println("No provera found ", err)
		return
	}
	if provera.Status != data.PUSTEN {
		http.Error(rw, "Provera not propusten", http.StatusNotAcceptable)
		p.logger.Println("Provera not propusten")
		return
	}
	prelazak := &data.PrelazakGranice{}
	prelazak.Vreme = time.Now().Format("02-01-06 15:04")
	prelazak.Provera = provera
	prelazak.GPrelaz = policajac.GPrelaz
	if p.repo.CreatePrelazak(prelazak) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}
	rw.WriteHeader(http.StatusInternalServerError)
}

func (p *PrelazakHandler) GetPrelasci(rw http.ResponseWriter, h *http.Request) {
	prelasci := p.repo.GetPrelasci()

	err := prelasci.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to Json", http.StatusInternalServerError)
		p.logger.Println("Unable to convert to Json: ", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *PrelazakHandler) GetPrelasciByGPrelaz(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var gprelaz = vars["g_prelaz"]

	prelasci := p.repo.GetPrelasciByPrelaz(gprelaz)

	err := prelasci.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to Json", http.StatusInternalServerError)
		p.logger.Println("Unable to convert to Json: ", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *PrelazakHandler) MiddlewarePrelazakValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		prelazak := &data.PrelazakGranice{}
		err := prelazak.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), PrelazakKey{}, prelazak)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
