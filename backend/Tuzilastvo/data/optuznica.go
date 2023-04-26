package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Optuznica struct {
	Id              string          `json:"id"`
	IdGradjanina    string          `json:"idGradjanina"`
	Aktivna         bool            `json:"aktivna"`
	KrivicnaPrijava KrivicnaPrijava `json:"krivicnaPrijava"`
	//Za dodati sta sve treba
}

type Optuznice []*Optuznica

func (p *Optuznice) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Optuznica) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Optuznica) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Optuznica) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
