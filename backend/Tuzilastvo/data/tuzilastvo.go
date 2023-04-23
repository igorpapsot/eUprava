package data

import (
	"encoding/json"
	"io"
)

type Tuzilastvo struct {
	Id       string `json:"id"`
	Naziv    string `json:"naziv"`
	Lokacija Mesto  `json:"lokacija"`
}

type Tuzilastva []*Tuzilastvo

func (p *Tuzilastva) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tuzilastvo) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tuzilastvo) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
