package main

import (
	//"fmt"
	"github.com/K-logeshwaran/goDb/Driver"
	"log"
	"os"
	// "github.com/K-logeshwaran/goDb/handlers"
)

/*
To do :
M
	last modified : 29/08/2023 12:37 AM
*/

func main() {
	l := log.New(os.Stdout, "myJSON DB reports -> ", log.LstdFlags)
	DB := Driver.NewDB("./database", l, Driver.NewCollection())
	DB.CreateDB()
	DB.CreateCollection("dev34")
	DB.PopulateRecords("Users", []byte(`
	{
		"dep": "BCA",
		"exp": 324,
		"spec":"web dev2",
		"name": "Tren"
	}
	`))
	val, _ := DB.Where("Users", "name", "Tren")
	println(len(val))
	for _, v := range val {
		println(v)
	}
	println(DB.ListCollections().ToJson())

	// fmt.Println(DB.IsCollectionExist("Users"))
	// w, _ := DB.UpdateOneById("Admin", "d07f2928-7ac0-4c15-89ae-d42beebb6860", "name22", "punda")
	// //DB.AddField("Admin", "d07f2928-7ac0-4c15-89ae-d42beebb6860", "name3", "Junni")

	// if w != nil {
	// 	fmt.Println(w.ToJson())
	// }
	// fmt.Println(w)

}
