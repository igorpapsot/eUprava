package data

import (
	"encoding/json"
	"io"
)

type Optuzeni struct {
	Ime               string `json:"ime"`
	Prezime           string `json:"prezime"`
	DatumRodjenja     string `json:"datumRodjenja"`
	Jmbg              string `json:"jmbg"`
	Zanimanje         string `json:"zanimanje"`
	BrTelefona        string `json:"brTelefona"`
	MestoPrebivalista Mesto  `json:"mestoPrebivalista"`
}

type Optuzenici []*Optuzeni

func (p *Optuzenici) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Optuzeni) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Optuzeni) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
