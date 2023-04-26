package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type KrivicnaPrijava struct {
	Id           string        `json:"id"`
	Privatnost   bool          `json:"privatnost"`
	ClanZakonika string        `json:"clanZakonika"`
	Datum        string        `json:"datum"`
	MestoPrijave string        `json:"mestoPrijave"`
	TuzilastvoId string        `json:"tuzilastvoId"`
	Obrazlozenje string        `json:"obrazlozenje"`
	Status       StatusPrijave `json:"status"`
	Optuzeni     Optuzeni      `json:"optuzeni"`
}

type KrivicnePrijave []*KrivicnaPrijava

func (p *KrivicnePrijave) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *KrivicnaPrijava) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *KrivicnaPrijava) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *KrivicnaPrijava) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
