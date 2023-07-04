package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Rociste struct {
	Id       string  `json:"id"`
	sudijaId string  `json:"sudijaId"`
	Datum    string  `json:"datum"`
	Mesto    string  `json:"mesto"`
	Sud      sudEnum `json:"sud"`
}

type Rocista []*Rociste

func (p *Rocista) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Rociste) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Rociste) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Rociste) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
