package handlers

import (
	"io"
	"log"

	"github.com/K-logeshwaran/goDb/Driver"

	//"io"
	"fmt"
	"net/http"
)

func TOBYTES(s string) []byte {
	return []byte(s)
}

type DBApi struct {
	D *Driver.DataBase
}

func (a *DBApi) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Logger value")
	log.Println("dassadasd")
	fmt.Println("Listening on dasdasda")
	rw.Write(TOBYTES("Bro am working"))
}

func NewApi(loc string, logger string, col Driver.Collection) *DBApi {
	return &DBApi{
		D: Driver.NewDB(loc, logger, col),
	}
}

func (a *DBApi) Collection(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		clc := r.FormValue("collection")
		err := a.D.CreateCollection(clc)
		if err != nil {
			log.Fatal(err)
		}
		rw.Write([]byte(clc))
	case http.MethodGet:
		clcObj := a.D.ListCollections()
		log.Println(clcObj.Value()["Collection"])
		rw.Write(TOBYTES("BRRRRRRRRRRRRRR"))
		rw.Write(clcObj.ToBytes())

	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}

func (a *DBApi) Records(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		clc := r.FormValue("collection")
		rw.Header().Set("Content-Type", "application/json")
		wraper, err := a.D.ReadAll(clc)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES("Collection Does not Exists "))
			return
		}
		rw.Write(Driver.WrapperArrayToBytes(wraper))

	case http.MethodPost:

		clc := r.FormValue("collection")
		rw.Header().Set("Content-Type", "application/json")
		post, _ := io.ReadAll(r.Body)
		a.D.PopulateRecords(clc, post)
	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}

func (a *DBApi) FindOne(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		id := r.FormValue("id")
		clc := r.FormValue("collection")
		rw.Header().Set("Content-Type", "application/json")
		wraper, _, err := a.D.FindOneById(clc, id)
		if err != nil {
			//log.Fatalln(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES("Record with " + id + "Doesnot exists"))
			return
		}
		rw.Write(wraper.ToBytes())
	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}
func (a *DBApi) Where(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		id := r.FormValue("field")
		value := r.FormValue("value")
		clc := r.FormValue("collection")
		rw.Header().Set("Cont	ent-Type", "application/json")
		wraper, err := a.D.Where(clc, id, value)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES("Record with " + id + "Doesnot exists"))
			return
		}
		rw.Write(Driver.WrapperArrayToBytes(wraper))
	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}
