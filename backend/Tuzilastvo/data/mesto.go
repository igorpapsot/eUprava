package data

import (
	"encoding/json"
	"io"
)

type Mesto struct {
	Id    string `json:"id"`
	Naziv string `json:"naziv"`
	Ulica string `json:"ulica"`
	Broj  int    `json:"broj"`
}

type Mesta []*Mesto

func (p *Mesta) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Mesto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Mesto) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
