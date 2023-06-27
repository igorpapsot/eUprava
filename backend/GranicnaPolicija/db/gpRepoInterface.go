package db

import "GranicnaPolicija/data"

type GpRepo interface {
	LoginUser(username string, password string) (data.GPolicajac, error)
	GetUserByJMBG(jmbg int) (data.Gradjanin, error)
}
