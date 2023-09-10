package main

import (
	"flag"
	"fmt"
	"github.com/K-logeshwaran/goDb/Driver"
	"github.com/K-logeshwaran/goDb/handlers"
	"log"
	"net/http"
)

func TOBYTES(s string) []byte {
	return []byte(s)
}

/*
=============================================================
Welcome to ToyDB - Your Lightweight JSON Database Solution
=============================================================

ToyDB is ready to help you manage your data efficiently. Here are a few things you can do:

1. Create a new database or use an existing one.
2. Organize your data into collections for easy access.
3. Insert, update, or delete records effortlessly using JSON documents.
4. Perform queries to retrieve the information you need.

To get started, use the command-line interface (CLI) with the following commands:

- Create a new database:
  $ ./ToyDB -location /path/to/your/database

- Start the server and interact with ToyDB:
  $ ./ToyDB -location /path/to/your/database -port [desired_port]

For help and more information, run:
$ ./ToyDB --help

Thank you for choosing ToyDB. Let's simplify your data management!

*/

func main() {

	var (
		dbLoc     string
		PORT      string
		runserver bool
	)
	flag.StringVar(&dbLoc, "location", "./database", "Location of your Database")
	flag.StringVar(&PORT, "port", "2080", "sets port for database api")
	flag.BoolVar(&runserver, "serve", false, "starts the server on the given port")
	flag.Parse()
	if !runserver {
		fmt.Println("\t=============================================================")
		fmt.Println("\t Welcome to ToyDB - Your Lightweight JSON Database Solution")
		fmt.Println("\t=============================================================")

		fmt.Print(`
ToyDB is ready to help you manage your data efficiently. Here are a few things you can do:

	1. Create a new database or use an existing one.
	2. Organize your data into collections for easy access.
	3. Insert, update, or delete records effortlessly using JSON documents.
	4. Perform queries to retrieve the information you need.
`)
		fmt.Print("\nAVAILABLE FLAGs:\n\n")
		flag.PrintDefaults()
		s := ""
		fmt.Scan(&s)
		return

	} else {

		logFileLoc := dbLoc + "/logger.log"

		api := handlers.NewApi(dbLoc, logFileLoc, Driver.NewCollection(dbLoc))
		fmt.Println("Listening on  http://localhost:" + PORT)
		log.Println("Listening on  http://localhost:" + PORT)
		mux := http.NewServeMux()
		mux.HandleFunc("/", api.ServeHTTP)
		mux.HandleFunc("/collection", api.Collection)
		mux.HandleFunc("/records", api.Records)
		mux.HandleFunc("/findone", api.FindOne)
		mux.HandleFunc("/where", api.Where)
		mux.HandleFunc("/update", api.Update)
		mux.HandleFunc("/addField", api.AddNewField)

		http.ListenAndServe(":"+PORT, mux)

	}
}
