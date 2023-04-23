package data

import (
	"encoding/json"
	"io"
)

type Dokaz struct {
	Id                string `json:"id"`
	Tekst             string `json:"tekst"`
	Datum             string `json:"datum"`
	KrivicnaPrijavaId int    `json:"krivicnaPrijavaId"`
}

type Dokazi []*Dokaz

func (p *Dokazi) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Dokaz) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Dokaz) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
