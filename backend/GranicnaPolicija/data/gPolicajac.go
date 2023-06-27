package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type GPolicajac struct {
	Id       string   `json:"id"`
	Ime      string   `json:"ime"`
	Prezime  string   `json:"prezime"`
	Jmbg     string   `json:"jmbg"`
	Lozinka  string   `json:"lozinka"`
	Pol      EPol     `json:"pol"`
	GPrelaz  EGPrelaz `json:"prelaz"`
	Password string   `json:"password"`
	CCode    int      `json:"ccode"`
}

type GPolicajci []*GPolicajac

func (p *GPolicajci) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *GPolicajac) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *GPolicajac) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *GPolicajac) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func (p *GPolicajac) HashPassword(lozinka string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(lozinka), 14)
	return string(bytes), err
}
