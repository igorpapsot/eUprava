package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

type ProveraHandler struct {
	logger *log.Logger
	repo   db.GpRepo
}

type ProveraKey struct{}

type PrijavaKey struct{}

func NewProveraHandler(l *log.Logger, ur db.GpRepo) *ProveraHandler {
	return &ProveraHandler{l, ur}
}

func (p ProveraHandler) CreateProveraHandler(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var gradjaninJmbg = vars["gradjanin"]
	var policajacId = vars["policajacId"]

	client := &http.Client{}
	req1, err := http.NewRequest("GET", "http://localhost:8004/user/jmbg", nil)
	if err != nil {
		//error
		return
	}
	req1.Header.Add("jmbg", gradjaninJmbg)
	resp1, err := client.Do(req1)
	req2, err := http.NewRequest("GET", "http://localhost:8004/poternica/gradjanin", nil)
	if err != nil {
		//error
		return
	}
	req2.Header.Add("jmbg", gradjaninJmbg)
	resp2, err := client.Do(req2)

	var provera *data.ProveraGradjanina
	provera.Vreme = time.Now().String()
	provera.Status = data.EStatusProvere(data.NACEKANJU)

	gradjaninJson, err := io.ReadAll(resp1.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to retrive gradjanin.", err)
		return
	}
	var gradjanin data.Gradjanin
	err1 := json.Unmarshal(gradjaninJson, &gradjanin)
	if err1 != nil {
		http.Error(rw, err1.Error(), http.StatusNotFound)
		p.logger.Println("Unable to unmarshal gradjanin.", err1)
		return
	}
	provera.Gradjanin = gradjanin

	policajac, err := p.repo.GetPolicajac(policajacId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to find policajac.", err)
		return
	}
	provera.Policajac = policajac

	poternicaJson, err := io.ReadAll(resp2.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to retrive poternica.", err)
		return
	}
	var poternica data.Poternica
	err2 := json.Unmarshal(poternicaJson, &poternica)
	if err2 != nil {
		http.Error(rw, err2.Error(), http.StatusNotFound)
		p.logger.Println("Unable to unmarshal poternica.", err2)
		return
	}
	provera.Poternica = poternica

	if p.repo.CreateProvera(provera) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)

}

// Verovatno ne treba jer cu slati direkt Tuzilastvu novu krivicnu prijavu
func (p ProveraHandler) Prijava(rw http.ResponseWriter, h *http.Request) {
	prijava := h.Context().Value(PrijavaKey{}).(*data.KrivicnaPrijava)

	if p.repo.CreatePrijava(prijava) {
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte("202 - Accepted"))
		return
	}
}

func (p *ProveraHandler) GetProvereNaCekanju(rw http.ResponseWriter, h *http.Request) {
	provere := p.repo.GetProvereByStatus(data.NACEKANJU)

	err := provere.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to Json", http.StatusInternalServerError)
		p.logger.Println("Unable to convert to Json: ", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *ProveraHandler) GetProvere(rw http.ResponseWriter, h *http.Request) {
	provere := p.repo.GetProvere()

	err := provere.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to convert to Json", http.StatusInternalServerError)
		p.logger.Println("Unable to convert to Json: ", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *ProveraHandler) MiddlewareProveraValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		provera := &data.ProveraGradjanina{}
		err := provera.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), ProveraKey{}, provera)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
