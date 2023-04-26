package db

import "Tuzilastvo/data"

type Repo interface {
	CreateTuzilastvo(p *data.Tuzilastvo) bool
	CreatePrijava(prijava *data.KrivicnaPrijava) bool
	ConfirmPrijava(id string) bool
	DeclinePrijava(id string) bool
	GetPrijave(javne bool) data.KrivicnePrijave
	GetPrijava(id string) (data.KrivicnaPrijava, error)
	GetTuzilastvo(id string) (data.Tuzilastvo, error)
	GetTuzilastva() data.Tuzilastva
	CreateOptuznica(optuznica *data.Optuznica) bool
	GetOptuznice() data.Optuznice
	GetOptuznica(prijavaId string) (data.Optuznica, error)
	//SendKrivicnaPrijava() ??
}
