package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Tuzilac struct {
	Id      string `json:"id"`
	Ime     string `json:"ime"`
	Prezime string `json:"prezime"`
	Pol     Pol    `json:"pol"`
	Jmbg    string `json:"jmbg"`
	Lozinka string `json:"lozinka"`
}

type Tuzioci []*Tuzilac

func (p *Tuzioci) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tuzilac) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tuzilac) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Tuzilac) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
