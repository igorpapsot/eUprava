package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type ProveraGradjanina struct {
	Id              string    `json:"id"`
	Gradjanin       Gradjanin `json:"gradjanin"`
	Vreme           string    `json:"vreme"`
	ZabranaPrelaska bool      `json:"zabrana_prelaska"`
	Poternica       Poternica `json:"poternica"`
}

type ProvereG []*ProveraGradjanina

func (p *ProvereG) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *ProveraGradjanina) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *ProveraGradjanina) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *ProveraGradjanina) ToBson(doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}