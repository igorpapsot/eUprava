package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Optuzeni struct {
	Ime     string   `json:"ime"`
	Prezime string   `json:"prezime"`
	Jmbg    string   `json:"jmbg"`
	Lozinka string   `json:"lozinka"`
	Pol     EPol     `json:"pol"`
	GPrelaz EGPrelaz `json:"prelaz"`
}

func (p *Optuzeni) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Optuzeni) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Optuzeni) ToBson(doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
