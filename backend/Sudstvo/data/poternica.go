package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Poternica struct {
	Id          string `json:"id"`
	sudijaId    string `json:"sudijaId"`
	gradjaninId string `json:"gradjaninId"`
	naslov      string `json:"naslov"`
	opis        string `json:"opis"`
}

type Poternice []*Poternica

func (p *Poternice) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Poternica) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Poternica) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Poternica) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
