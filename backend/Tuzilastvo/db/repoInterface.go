package db

import "Tuzilastvo/data"

type Repo interface {
	CreatePrijava(prijava *data.KrivicnaPrijava) bool
	ConfirmPrijava(prijava *data.KrivicnaPrijava) bool
	DeclinePrijava(prijava *data.KrivicnaPrijava) bool
	GetPrijave() data.KrivicnePrijave
	//SendKrivicnaPrijava() ??
}
