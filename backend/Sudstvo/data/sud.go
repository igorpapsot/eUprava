package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Sud struct {
	Id    string  `json:"id"`
	Sud   sudEnum `json:"sud"`
	Datum string  `json:"datum"`
	Mesto string  `json:"mesto"`
}

type Sudovi []*Sud

func (p *Sudovi) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Sud) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Sud) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Sud) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
