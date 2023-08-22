package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"context"
	"encoding/json"
	"fmt"
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

	//TODO: proveri jel radi ovo
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
	provera.Vreme = time.Now().String()
	provera.Status = data.EStatusProvere(data.NACEKANJU)

	fmt.Println(provera)

	//if p.repo.CreateProvera(provera) {
	//	rw.WriteHeader(http.StatusAccepted)
	//	return
	//}
	err = provera.ToJSON(rw)
	rw.WriteHeader(http.StatusCreated)

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
