package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type Gradjanin struct {
	Id      string `json:"id"`
	Ime     string `json:"ime"`
	Prezime string `json:"prezime"`
	Jmbg    string `json:"jmbg"`
	Pol     EPol   `json:"pol"`
}

func (p *Gradjanin) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Gradjanin) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Gradjanin) ToBson(doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
