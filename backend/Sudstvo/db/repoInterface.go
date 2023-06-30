package db

import "Sudstvo/data"

type Repo interface {
	PostOptuznica(optuznica *data.Optuznica) bool
	GetOptuznice() data.Optuznice
	GetOptuznica(id string) (data.Optuznica, error)
	CreatePoternica(poternica *data.Poternica) bool
	GetPoternice() data.Poternice
	GetPoternica(id string) (data.Poternica, error)
	CreateKonacnaPresuda(p *data.KonacnaPresuda) bool         //ispraviti
	GetKonacnaPresuda(id string) (data.KonacnaPresuda, error) //ispraviti verovatno
	CreateSud(p *data.Sud) bool
	GetSud(id string) (data.Sud, error)
	GetSudovi() data.Sudovi
	CreateSudija(p *data.Sudija) bool
	GetSudija(jmbg string) (data.Sudija, error)
	GetSudije() data.Sudije
	CreateRociste(rociste *data.Rociste) bool
	GetRocista() data.Rocista
	GetRociste(id string) (data.Rociste, error)
	Login(jmbg string, lozinka string) (data.Sudija, error)
	Register(sudija *data.Sudija) bool
}
