package handlers

import (
	"github.com/K-logeshwaran/goDb/Driver"
	//"log"
	//"io"
	"net/http"
)

func TOBYTES(s string) []byte {
	return []byte(s)
}

type DBApi struct {
	D *Driver.DataBase
}

func (a *DBApi) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	a.D.Logger.Println()
	rw.Write(TOBYTES("Bro am working"))
}

func NewApi(loc string, logger string, col Driver.Collection) *DBApi {
	return &DBApi{
		D: Driver.NewDB(loc, logger, col),
	}
}
