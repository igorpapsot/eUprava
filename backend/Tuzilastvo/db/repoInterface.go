package db

import "Tuzilastvo/data"

type Repo interface {
	CreateTuzilastvo(p *data.Tuzilastvo) bool
	CreatePrijava(prijava *data.KrivicnaPrijava) bool
	ConfirmPrijava(prijava *data.KrivicnaPrijava) bool
	DeclinePrijava(prijava *data.KrivicnaPrijava) bool
	GetPrijave() data.KrivicnePrijave
	GetPrijava(id string) (data.KrivicnaPrijava, error)
	GetTuzilastvo(id string) (data.Tuzilastvo, error)
	GetTuzilastva() data.Tuzilastva
	//SendKrivicnaPrijava() ??
}
