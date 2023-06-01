package data

import (
	"encoding/json"
	"io"
)

type Stranka struct {
	Id          string `json:"id"`
	brojClanova int    `json:"brojClanova"`
}

type Stranke []*Stranka

func (p *Stranke) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Stranka) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Stranka) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
