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

func CreateCollectionFiles(loc string) {
	fp := path.Join(loc, "collections.json")
	_, err := os.Stat(fp)
	log.Println(fp)
	if os.IsNotExist(err) {
		fs, _ := os.Create(fp)
		fs.Close()
		log.Println("Collection created success")
	}

}

func NewCollection(dbloc string) Collection {
	d, _ := os.ReadFile(dbloc + "/collections.json")
	c := Collection{}
	json.Unmarshal(d, &c)
	return c
}

func (c *Collection) AddCollection(cl string) {
	c.Collections = append(c.Collections, cl)
}

func (c *Collection) Commit(dbloc string) {
	data, _ := json.Marshal(c)

	err := os.WriteFile(dbloc+"/collections.json", data, 0644)
	//fs, err := os.OpenFile(dbloc+"/collections.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//fs.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
