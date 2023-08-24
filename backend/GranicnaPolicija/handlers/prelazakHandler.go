package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"context"
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

func NewPrelazakHandler(l *log.Logger, ur db.GpRepo) *PrelazakHandler {
	return &PrelazakHandler{l, ur}
}

// TODO : dodati granicni prelaz (preko policajca izvuci info)
// IDEA : ako je provera prosao smes da napravis prelazak inace baci error
func (p *PrelazakHandler) CreatePrelazakHandler(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var proveraId = vars["proveraId"]
	provera, err := p.repo.GetProveraById(proveraId)
	if err != nil {
		http.Error(rw, "No provera found", http.StatusNotAcceptable)
		p.logger.Println("No provera found ", err)
		return
	}
	prelazak := &data.PrelazakGranice{}
	prelazak.Vreme = time.Now().String()
	prelazak.Provera = provera

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
