package Driver

import (
	"errors"
	"fmt"
	"strings"

	"log"
	"os"

	"github.com/google/uuid"

	"path"
)

type DataBase struct {
	Location    string
	Logger      *log.Logger
	collections Collection
	//FileChan    chan []byte
}

// Done
func doesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)
	return os.IsNotExist(error)
}
func NewDB(loc string, logger string, col Collection) *DataBase {
	_, err := os.Stat(loc)
	if os.IsNotExist(err) {
		err = os.Mkdir(loc, 0777)
		if err != nil {
			panic(err)
		}
	}
	var (
		fs *os.File
		e  error
	)
	if doesFileExist(logger) {
		fs, e = os.Create(logger)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("Created")
	} else {
		fs, _ = os.Open(logger)
		fmt.Println("alredy")
	}
	l := log.New(fs, "myJSON DB reports -> ", log.LstdFlags)

	CreateCollectionFiles("."+loc, l)

	return &DataBase{
		Location: loc,
		Logger:   l,
		//FileChan: c,
		collections: col,
	}
}

// Done

// Done
func (d *DataBase) CreateCollection(name string) {
	loc := path.Join(d.Location, name)

	_, err := os.Stat(loc)

	if os.IsNotExist(err) {
		err = os.Mkdir(loc, 0777)
		d.collections.AddCollection(name)
		d.Logger.Println(d.collections)
		d.collections.Commit()
		if err != nil {
			d.Logger.Fatal(err)
		}
	} else {
		d.Logger.Printf("%s collection already exists \n", name)
	}
}

func (d *DataBase) IsCollectionExist(name string) bool {
	loc := path.Join(d.Location, name)
	_, err := os.Stat(loc)
	return !os.IsNotExist(err)
}

// Done
func (d *DataBase) PopulateRecords(collection string, data []byte) {

	ObjId := createuuid()
	fileName := ObjId + ".json"
	fileLocation := path.Join(d.Location, collection, fileName)
	jsonMap := BuildWrapper(data)
	jsonMap.AddField("id", ObjId)

	file, err := os.Create(fileLocation)

	if err != nil {
		d.Logger.Fatal(err)
	}
	_, err = file.Write(jsonMap.ToBytes())

	if err != nil {
		d.Logger.Fatal(err)
	}

	d.Logger.Println("Data Addes successfully")

}

// Done
func createuuid() string {
	return uuid.New().String()
}

// Done
func (d *DataBase) ReadAll(collection string) []Wrapper {
	w := []Wrapper{}
	loc := path.Join(d.Location, collection)
	records, err := os.ReadDir(loc)
	if err != nil {
		panic("Dir not found")
	}
	for _, record := range records {

		r, err := os.ReadFile(path.Join(d.Location, collection, record.Name()))
		if err != nil {
			panic("something went wrong line 98")
		}
		w = append(w, *BuildWrapper(r))
	}
	return w

}

// Done
func (d *DataBase) FindOneById(collection string, id string) (*Wrapper, string, error) {
	if d.IsCollectionExist(collection) {
		loc := path.Join(d.Location, collection)
		files, err := os.ReadDir(loc)
		if err != nil {
			panic("Err line 113")
		}
		for _, v := range files {
			if strings.Split(v.Name(), ".")[0] == id {
				d, err := os.ReadFile(path.Join(loc, v.Name()))
				if err != nil {
					panic("Error 119")
				}
				//return *BuildWrapper(d), nil
				wrapper := BuildWrapper(d)
				return wrapper, path.Join(loc, v.Name()), nil
			}
		}
		return nil, "", errors.New("no record found")
	} else {
		return nil, "", errors.New("collection not found")
	}
}

func (d *DataBase) UpdateOneById(collection, id, filed string, value interface{}) (*Wrapper, error) {
	w, filePath, err := d.FindOneById(collection, id)
	if err != nil {
		return nil, err
	}
	if w.Value()[filed] == nil {
		return nil, nil
	}
	w.AddField(filed, value)
	d.commit(filePath, w)
	return w, nil
}

func (d *DataBase) AddField(collection, id, filed string, value interface{}) (*Wrapper, error) {
	w, filePath, err := d.FindOneById(collection, id)
	if err != nil {
		return nil, err
	}

	w.AddField(filed, value)
	d.commit(filePath, w)
	return w, nil
}

func (d *DataBase) commit(recordpath string, w *Wrapper) {
	os.WriteFile(recordpath, w.ToBytes(), os.ModeAppend)
}

func (d *DataBase) ListCollections() *Wrapper {

	da, _ := os.ReadFile(COLLECTIONFILESLOC)
	return BuildWrapper(da)
}

func (d *DataBase) Where(collection string, field string, value interface{}) ([]string, error) {
	var reA []string
	if d.IsCollectionExist(collection) {
		loc := path.Join(d.Location, collection)
		files, err := os.ReadDir(loc)
		if err != nil {
			panic("Err line 113")
		}
		for _, v := range files {
			d, _ := os.ReadFile(path.Join(loc, v.Name()))
			w := BuildWrapper(d)
			if w.data[field] == value {
				reA = append(reA, w.ToJson())
			}
		}
		return reA, nil
	} else {
		return nil, errors.New("collection not found")
	}

}
