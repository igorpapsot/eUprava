package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"context"
	"encoding/json"
	"fmt"
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

type ProveraStruct struct {
	PolicajacId string //`json:"policajacId"`
	Gradjanin   string //`json:"gradjanin"`
}

type ProveraKey struct{}

type PrijavaKey struct{}

func NewProveraHandler(l *log.Logger, ur db.GpRepo) *ProveraHandler {
	return &ProveraHandler{l, ur}
}

func (p ProveraHandler) CreateProveraHandler(rw http.ResponseWriter, h *http.Request) {

	decoder := json.NewDecoder(h.Body)
	var proveravanje ProveraStruct
	err := decoder.Decode(&proveravanje)

	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusInternalServerError)
		p.logger.Println("Unable to decode json :", err)
		return
	}
	p.logger.Println(proveravanje)

	policajac, err := p.repo.GetPolicajac(proveravanje.PolicajacId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to find policajac.", err)
		return
	}
	fmt.Println(policajac.Ime)

	client := &http.Client{}
	req1, err := http.NewRequest("GET", "http://mup_service:8004/user/jmbg", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	q := req1.URL.Query()
	q.Add("jmbg", proveravanje.Gradjanin)
	req1.URL.RawQuery = q.Encode()
	resp1, err := client.Do(req1)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		fmt.Println("zahtev zeznuo")
		return
	}

	gradjaninJson, err := io.ReadAll(resp1.Body)
	fmt.Println(gradjaninJson)
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

	if gradjanin.Id == "" {
		http.Error(rw, "gradjanin not found", http.StatusNotFound)
		p.logger.Println("gradjanin not found", err1)
		return
	}

	req2, err := http.NewRequest("GET", "http://mup_service:8004/poternica/gradjanin", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	q2 := req2.URL.Query()
	q2.Add("gradjaninId", gradjanin.Id)
	req2.URL.RawQuery = q2.Encode()
	resp2, err := client.Do(req2)

	poternicaJson, err := io.ReadAll(resp2.Body)
	if err != nil {
		//http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to retrive poternica.", err)
		rw.Write([]byte("nema poternicu"))
		return
	}

	var poternica data.Poternica
	err2 := json.Unmarshal(poternicaJson, &poternica)
	if err2 != nil {
		http.Error(rw, err2.Error(), http.StatusNotFound)
		p.logger.Println("Unable to unmarshal poternica.", err2)
		return
	}
	fmt.Println(poternicaJson)
	fmt.Println("dosao do kreiranja")

	provera := &data.ProveraGradjanina{}
	provera.Policajac = policajac
	provera.Gradjanin = gradjanin
	provera.Poternica = poternica
	provera.Vreme = time.Now().Format("02-01-06 15:04")
	provera.Status = data.EStatusProvere(data.NACEKANJU)

	fmt.Println(provera)

	if p.repo.CreateProvera(provera) {
		provera.ToJSON(rw)
		rw.WriteHeader(http.StatusAccepted)
		return
	}
	rw.WriteHeader(http.StatusInternalServerError)

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

func (p *ProveraHandler) RefusePrelazak(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var proveraId = vars["proveraId"]
	provera, err := p.repo.GetProveraById(proveraId)
	if err != nil {
		http.Error(rw, "No provera found", http.StatusNotAcceptable)
		p.logger.Println("No provera found ", err)
		return
	}

	provera.Status = data.ODBIJEN

	if !p.repo.UpdateProvera(&provera) {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
}

func (p *ProveraHandler) AcceptPrelazak(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var proveraId = vars["proveraId"]
	provera, err := p.repo.GetProveraById(proveraId)
	if err != nil {
		http.Error(rw, "No provera found", http.StatusNotAcceptable)
		p.logger.Println("No provera found ", err)
		return
	}

	provera.Status = data.PUSTEN

	if !p.repo.UpdateProvera(&provera) {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
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

//TODO : get prijave na cekanju da bih se preko fronta poslale tuzilastvu
//TODO : takodje napraviti methodu za to u mainu NE ZABORAVI JEEZ OSNOVNO

// Verovatno ne treba jer cu slati direkt Tuzilastvu novu krivicnu prijavu
func (p ProveraHandler) Prijava(rw http.ResponseWriter, h *http.Request) {
	prijava := h.Context().Value(PrijavaKey{}).(*data.KrivicnaPrijava)

	prijava.Status = data.EStatusPrijave(data.CEKANASLANJE)

	if p.repo.CreatePrijava(prijava) {
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte("202 - Accepted"))
		return
	}

	rw.WriteHeader(http.StatusInternalServerError)
}

func (p ProveraHandler) PoslataPrijava(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var prijavaId = vars["prijavaId"]
	prijava, err := p.repo.GetPrijava(prijavaId)
	if err != nil {
		http.Error(rw, "No prijava found", http.StatusNotAcceptable)
		p.logger.Println("No prijava found ", err)
		return
	}

	prijava.Status = data.EStatusPrijave(data.PROSLEDJENO)

	if !p.repo.UpdatePrijava(&prijava) {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
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
