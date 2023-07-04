package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type KonacnaPresuda struct {
	Id      string `json:"id"`
	Aktivna bool   `json:"aktivna"`
	Opis    string `json:"opis"`
}

type KonacnePresude []*KonacnaPresuda

func (p *KonacnePresude) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *KonacnaPresuda) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *KonacnaPresuda) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *KonacnaPresuda) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
