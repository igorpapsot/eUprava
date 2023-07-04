package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Sudija struct {
	Id      string  `json:"id"`
	Ime     string  `json:"ime"`
	Prezime string  `json:"prezime"`
	Pol     polEnum `json:"pol"`
	Jmbg    string  `json:"jmbg"`
	Lozinka string  `json:"lozinka"`
	Sud     sudEnum `json:"sud"`
}

type Sudije []*Sudija

func (p *Sudije) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Sudija) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Sudija) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Sudija) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
