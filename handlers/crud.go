package handlers

import (
	"io"

	//"github.com/Jeffail/gabs"

	"github.com/K-logeshwaran/goDb/Driver"
	//"io"
	"net/http"
)

type Handler struct {
	db *Driver.DataBase
}

func (h *Handler) addRecords(c string, d []byte) {
	h.db.PopulateRecords(c, d)
}

func (h *Handler) readRecords() string {
	return h.db.ReadAll("Users")
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		R := h.readRecords()
		res.Header().Set("Content-Type", "application/json")
		c, err := gabs.ParseJSON([]byte(R))
		if err != nil {
			panic(err)
		}
		res.Write(c.Bytes())

	case http.MethodPost:
		collection := req.FormValue("collection")
		b, err := io.ReadAll(req.Body)
		h.db.Logger.Println(b)
		if err != nil {
			h.db.Logger.Fatalln(err)
		}
		h.addRecords(collection, b)
	}

}

func NewHandler(db *Driver.DataBase) *Handler {
	return &Handler{
		db: db,
	}
}
