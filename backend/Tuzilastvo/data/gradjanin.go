package data

import (
	"encoding/json"
	"io"
)

type Gradjanin struct {
	Id            string `json:"id"`
	Ime           string `json:"ime"`
	Prezime       string `json:"prezime"`
	Jmbg          string `json:"jmbg"`
	Pol           Pol    `json:"pol"`
	DatumRodjenja string `json:"datumRodjenja"`
	Osudjen       bool   `json:"osudjen"`
	Drzavljanstvo string `json:"drzavljanstvo"`
	MestoRodjenja string `json:"mestoRodjenja"`
}

type Gradjani []*Gradjanin

func (p *Gradjani) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Gradjanin) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Gradjanin) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
