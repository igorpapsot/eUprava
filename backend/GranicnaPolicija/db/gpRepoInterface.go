package db

import "GranicnaPolicija/data"

type GpRepo interface {
	Login(jmbg string, password string) (data.GPolicajac, error)
	NewGPolicajac(gpolicajac *data.GPolicajac) bool
	GetPolicajac(id string) (data.GPolicajac, error)

	CreatePrelazak(prelazak *data.PrelazakGranice) bool
	GetPrelasci() data.PrelasciGranice
	GetPrelazak(id string) (data.PrelazakGranice, error)

	CreatePrijava(prijava *data.KrivicnaPrijava) bool
	GetPrijave() data.KrivicnePrijave
	GetPrijava(id string) (data.KrivicnaPrijava, error)
}
