package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type PrelazakHandler struct {
	logger *log.Logger
	repo   db.GpRepo
}

type PrelazakKey struct{}

func (p *PrelazakHandler) NewPrelazakHandler(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var provera = vars["proveraId"]
	var prelazak *data.PrelazakGranice
	prelazak.Vreme = time.Now().String()
	prelazak.ProveraId = provera

	if p.repo.CreatePrelazak(prelazak) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
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
