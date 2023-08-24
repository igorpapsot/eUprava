package db

import "GranicnaPolicija/data"

type GpRepo interface {
	Login(jmbg string, password string) (data.GPolicajac, error)
	NewGPolicajac(gpolicajac *data.GPolicajac) bool
	GetPolicajac(id string) (data.GPolicajac, error)

	CreateProvera(provera *data.ProveraGradjanina) bool
	GetProvera(gradjanin *data.Gradjanin) (data.ProveraGradjanina, error)
	GetProveraById(id string) (data.ProveraGradjanina, error)
	GetProvere() data.ProvereG
	GetProvereByStatus(status string) data.ProvereG
	UpdateProvera(provera *data.ProveraGradjanina) bool

	CreatePrelazak(prelazak *data.PrelazakGranice) bool
	GetPrelasci() data.PrelasciGranice
	GetPrelasciByPrelaz(prelaz string) data.PrelasciGranice
	GetPrelazak(id string) (data.PrelazakGranice, error)

	CreatePrijava(prijava *data.KrivicnaPrijava) bool
	GetPrijave() data.KrivicnePrijave
	GetPrijava(id string) (data.KrivicnaPrijava, error)
	UpdatePrijava(prijava *data.KrivicnaPrijava) bool
}
