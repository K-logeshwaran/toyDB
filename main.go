// package main

// import (
// 	//"fmt"
// 	"github.com/K-logeshwaran/goDb/Driver"
// 	"log"
// 	"os"
// 	// "github.com/K-logeshwaran/goDb/handlers"
// )

// /*
// To do :
// M
// 	last modified : 29/08/2023 12:37 AM
// */

// func main() {
// 	l := log.New(os.Stdout, "myJSON DB reports -> ", log.LstdFlags)
// 	DB := Driver.NewDB("./database", l, Driver.NewCollection())
// 	DB.CreateDB()
// 	DB.CreateCollection("dev34")
// 	DB.PopulateRecords("Users", []byte(`
// 	{
// 		"dep": "BCA",
// 		"exp": 324,
// 		"spec":"web dev2",
// 		"name": "Tren"
// 	}
// 	`))
// 	val, _ := DB.Where("Users", "name", "Tren")
// 	println(len(val))
// 	for _, v := range val {
// 		println(v)
// 	}
// 	println(DB.ListCollections().ToJson())

// 	// fmt.Println(DB.IsCollectionExist("Users"))
// 	// w, _ := DB.UpdateOneById("Admin", "d07f2928-7ac0-4c15-89ae-d42beebb6860", "name22", "punda")
// 	// //DB.AddField("Admin", "d07f2928-7ac0-4c15-89ae-d42beebb6860", "name3", "Junni")

// 	// if w != nil {
// 	// 	fmt.Println(w.ToJson())
// 	// }
// 	// fmt.Println(w)

// }

package main

import (
	//"encoding/json"

	// //"api.com/my_api/routes"
	"flag"
	"fmt"
	"net/http"

	"github.com/K-logeshwaran/goDb/Driver"
	"github.com/K-logeshwaran/goDb/handlers"
)

func TOBYTES(s string) []byte {
	return []byte(s)
}

func main() {
	var (
		dbLoc string
		//logFileLoc string

	)
	flag.StringVar(&dbLoc, "location", "./database", "Location of your Database")
	//flag.StringVar(&logFileLoc, "logger", "logger.log", "Location of your Database Log file")
	flag.Parse()
	logFileLoc := dbLoc + "/logger.log"
	//	fmt.Println(doesFileExist(dbLoc + "/logger.log"))

	//l := log.New(fs, "myJSON DB reports -> ", log.LstdFlags)
	api := handlers.NewApi(dbLoc, logFileLoc, Driver.NewCollection())

	//DB := Driver.NewDB(dbLoc, l, Driver.NewCollection())
	// DB.CreateDB()
	// DB.CreateCollection("dev34")
	fmt.Println("Listening on  http://localhost:2080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.ServeHTTP)

	http.ListenAndServe(":2080", mux)

}
