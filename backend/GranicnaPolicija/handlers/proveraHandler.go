package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type ProveraHandler struct {
	logger *log.Logger
	repo   db.GpRepo
}

type ProveraKey struct{}

func (p ProveraHandler) NewProveraHandler(rw http.ResponseWriter, h *http.Request) {
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

	gradjaninJson := resp1.Body
	poternicaJson := resp2.Body

	var gradjanin data.Gradjanin
	err := json.Unmarshal(gradjaninJson, &gradjanin)
	pot := data.Poternica{}
	poternica := pot.FromJSON(poternicaJson)
	policajac, err := p.repo.GetPolicajac(policajacId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		p.logger.Println("Unable to find policajac.", err)
		return
	}

	var provera *data.ProveraGradjanina
	provera.Vreme = time.Now().String()
	provera.Policajac = policajac
	provera.Gradjanin = gradjanin

}

func (p ProveraHandler) PrijavaPoternice(rw http.ResponseWriter, h *http.Request) {

}

func (p ProveraHandler) PrijavaKrijumcarenja(rw http.ResponseWriter, h *http.Request) {

}
