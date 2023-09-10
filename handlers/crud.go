package handlers

import (
	"encoding/json"
	"io"
	"log"
	"strconv"

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
	rw.Write(TOBYTES(`
	<!DOCTYPE html>
<html>

<head>
    <title>ToyDB Database</title>
    <style>
        body {
            font-family: Arial, Helvetica, sans-serif;
            margin: 20px;
        }

        h1 {
            color: #333;
        }

        h2 {
            color: #444;
        }

        h3 {
            color: #555;
        }

        p {
            color: #666;
        }

        ol,
        ul {
            color: #777;
            margin-left: 20px;
        }

        pre {
            background-color: #f5f5f5;
            padding: 10px;
            border-radius: 5px;
        }

        a {
            color: #007bff;
            text-decoration: none;
        }

        hr {
            border: 1px solid #ddd;
        }

        code {
            background-color: #f8f8f8;
            border: 1px solid #ddd;
            border-radius: 3px;
            padding: 2px 5px;
        }
		#desc{
			text-align: center;
		}
    </style>
</head>

<body>

    <h1><a href="https://github.com/K-logeshwaran/toyDB">ToyDB Database<a></h1>

    <p>ToyDB is a lightweight JSON database designed to simplify data storage and retrieval for small to medium-scale applications.</p>

    <h2>Features</h2>

    <h3>JSON-Based</h3>
    <p>ToyDB uses JSON as the primary data format, making it easy to work with structured data. JSON is human-readable, making it straightforward to inspect and manipulate data.</p>

    <h3>File System Storage</h3>
    <p>Data is stored directly on the file system, eliminating the need for a separate database server. This approach simplifies data management and reduces overhead.</p>

    <h3>Collections</h3>
    <p>ToyDB organizes data into collections, allowing you to group related data together. Each collection acts as a container for JSON documents.</p>

    <h3>Query Capabilities</h3>
    <p>You can perform queries on your data using simple and intuitive commands, enabling efficient data retrieval.</p>

    <h3>Configurable</h3>
    <p>ToyDB is highly configurable, allowing you to specify the location of your database and customize various settings to meet your application's requirements.</p>

    <h2>Getting Started</h2>

    <p>To get started with ToyDB, follow these steps:</p>

    <ol>
        <li>Clone this repository to your local machine.</li>
        <li>Install ToyDB by running the installation script or following the installation instructions in the documentation.</li>
        <li>Create a new database or use an existing one.</li>
        <li>Use the command-line interface (CLI) to manage and query your data.</li>
    </ol>

    <p>For detailed usage instructions, refer to the documentation in the repository.</p>

    <h2>Usage</h2>

    <h3>Flags</h3>

    <ul>
        <li><code>-location</code>: Specifies the location of your database. Default value is <code>./database</code>.</li>
        <li><code>-port</code>: Sets the port for the database API. Default value is <code>2080</code>.</li>
        <li><code>-serve</code>: Starts the server on the given port if provided.</li>
    </ul>

    <h3>Running the Server</h3>

    <p>To start the ToyDB server, use the <code>-serve</code> flag. For example:</p>

    <pre>
        <code>./toydb -serve</code>
    </pre>

    <p>By default, the server will listen on <a href="http://localhost:2080">http://localhost:2080</a>.</p>

    <h2>API Endpoints</h2>

    <ul>
        <li><code>/: Welcome page with information about ToyDB</code></li>
        <li><code>/collection: Endpoint for managing collections.</code></li>
        <li><code>/records: Endpoint for managing records.</code></li>
        <li><code>/findone: Endpoint for finding a specific record.</code></li>
        <li><code>/where: Endpoint for querying records based on criteria.</code></li>
        <li><code>/update: Endpoint for updating records.</code></li>
        <li><code>/addField: Endpoint for adding a new field to a record.</code></li>
    </ul>

    <h2>Logging</h2>

    <p>ToyDB logs its activities to a file located at <code>{dbLoc}/logger.log</code> file.</p>

    <h2>Example</h2>

    <p>Here's an example of how to start the ToyDB server:</p>

    <pre>
        <code>./your-program -serve -location ./your-database-location -port 8080</code>
    </pre>

    <p>This will start the server on <a href="http://localhost:8080">http://localhost:8080</a>.</p>

    <p>Feel free to explore ToyDB and use it to manage your data efficiently.</p>

    <h2>Documentation</h2>

    <p>For more information on how to use ToyDB, consult the official documentation <a href="link-to-your-documentation">here</a>.</p>

    <h2>Contributing</h2>

    <p>We welcome contributions from the community! If you have suggestions, bug reports, or want to contribute code, please open an issue or submit a pull request on this repository.</p>

    <h2>License</h2>

    <p>ToyDB is open-source and released under the <a href="#">[License Name]</a> license. See the <a href="LICENSE">LICENSE</a> file for details.</p>

    <h2>Acknowledgments</h2>

    <p>We would like to thank the open-source community for their support and contributions to ToyDB.</p>

    <hr>

    <p id="desc">Thank you for choosing ToyDB. Simplify your data management with ease!</p>

</body>

</html>

	`))
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
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(clcObj.ToBytes())

	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}

func (a *DBApi) Records(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		clc := r.FormValue("collection")
		limit := r.FormValue("limit")
		done := make(chan bool, 1)
		reschan := make(chan Driver.Wrapper, 100)
		l, err := strconv.Atoi(limit)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(TOBYTES("InternalServerError"))
			log.Panic(err)
			return

		}
		wrapperArray := []Driver.Wrapper{}

		err = a.D.ReadAllGPt(clc, l, done, reschan)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES("Collection Does not Exists "))
			return
		}
		for wp := range reschan {
			fmt.Println("reading........")
			wrapperArray = append(wrapperArray, wp)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(Driver.WrapperArrayToBytes(wrapperArray))

		// for {
		// 	select {
		// 	case rc := <-wraper:
		// 		wrapperArray = append(wrapperArray, rc)
		// 	case d := <-done:
		// 		if d {
		// 			rw.Write(Driver.WrapperArrayToBytes(wrapperArray))
		// 		}
		// 	}
		// }

	case http.MethodPost:
		clc := r.FormValue("collection")
		//rw.Header().Set("Content-Type", "application/json")
		resultCh := make(chan string)
		errorCh := make(chan error)
		post, _ := io.ReadAll(r.Body)
		go func(c string, d []byte) {
			msg, err := a.D.PopulateRecords(c, d)
			if err != nil {
				errorCh <- err
			} else {
				resultCh <- msg
			}
		}(clc, post)
		select {
		case r := <-resultCh:
			rw.Write(TOBYTES("<h1>" + r + "</h1>"))
		case e := <-errorCh:
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES(e.Error()))
		}

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
		log.Println(id, value, clc)
		fmt.Println(id, value, clc)
		wraper, err := a.D.Where(clc, id, value)
		//users&field=name&value=Cynthia

		fmt.Println(err)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES("Record with " + id + "Doesnot exists"))
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(Driver.WrapperArrayToBytes(wraper))
	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}

func (a *DBApi) Update(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPut:
		field := r.FormValue("field")
		id := r.FormValue("id")
		value := r.FormValue("value")
		clc := r.FormValue("collection")
		w, err := a.D.UpdateOneById(clc, id, field, value)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES(err.Error()))
			return
		} else if w == nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES("No field  " + field + "found in the record found "))
			return
		} else {
			res := make(map[string]interface{}, 2)
			res["message"] = "Data Updated success fully"
			res["data"] = w.Value()
			d, err := json.Marshal(res)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write(TOBYTES("InternalServerError"))
				return
			}
			rw.Write(d)
		}
	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}

func (a *DBApi) AddNewField(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPut:
		field := r.FormValue("field")
		id := r.FormValue("id")
		value := r.FormValue("value")
		clc := r.FormValue("collection")
		w, err := a.D.AddField(clc, id, field, value)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(TOBYTES(err.Error()))
			return
		}
		res := make(map[string]interface{}, 2)
		res["message"] = "Data Updated success fully"
		res["data"] = w.Value()
		d, err := json.Marshal(res)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(TOBYTES("InternalServerError"))
			return
		}
		rw.Write(d)

	default:
		rw.Write(TOBYTES("METHOD NOT ALLOWED"))
	}
}
