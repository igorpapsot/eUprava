package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type PrelazakGranice struct {
	Id        string    `json:"id"`
	Gradjanin Gradjanin `json:"gradjanin"`
	Vreme     string    `json:"vreme"`
	GPrelaz   EGPrelaz  `json:"g_prelaz"`
}

type PrelasciGranice []*GPolicajac

func (p *PrelasciGranice) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *PrelazakGranice) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *PrelazakGranice) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *PrelazakGranice) ToBson(doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
