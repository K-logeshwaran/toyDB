package Driver

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type Collection struct {
	Collections []string `json:"Collections"`
}

const COLLECTIONFILESLOC string = "E:\\SideProjects\\myDataBase\\database\\collections.json"

func CreateCollectionFiles(loc string, log *log.Logger) {
	fp := path.Join(loc, "collections.json")
	_, err := os.Stat(fp)
	log.Println(fp)
	if os.IsNotExist(err) {
		fs, _ := os.Create(fp)
		fs.Close()
		log.Println("Collection created success")
	}

}

func NewCollection() Collection {
	d, _ := os.ReadFile("E:\\SideProjects\\myDataBase\\database\\collections.json")
	c := Collection{}
	json.Unmarshal(d, &c)
	return c
}

func (c *Collection) AddCollection(cl string) {
	c.Collections = append(c.Collections, cl)
}

func (c *Collection) Commit() {
	data, _ := json.Marshal(c)
	err := os.WriteFile("E:\\SideProjects\\myDataBase\\database\\collections.json", data, os.ModeAppend)
	if err != nil {
		panic(err)
	}
}
