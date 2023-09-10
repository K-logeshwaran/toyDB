# Handlers Package Documentation

The `handlers` package provides HTTP request handlers for interacting with a ToyDB through the `Driver` package.

## Functions

### `TOBYTES(s string) []byte`

Converts a string to a byte slice.

- `s` (string): The input string to be converted.
- Returns: A byte slice containing the converted string.

### `NewApi(loc string, logger string, col Driver.Collection) *DBApi`

Creates and returns a new `DBApi` instance for handling database requests.

- `loc` (string): The location of the database.
- `logger` (string): Logger information.
- `col` (Driver.Collection): A collection instance.
- Returns: A pointer to the `DBApi` instance.

### `(a *DBApi) ServeHTTP(rw http.ResponseWriter, r *http.Request)`

Handles the main HTTP requests for the database API.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

### `(a *DBApi) Collection(rw http.ResponseWriter, r *http.Request)`

Handles requests related to database collections.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

### `(a *DBApi) Records(rw http.ResponseWriter, r *http.Request)`

Handles requests related to database records.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

### `(a *DBApi) FindOne(rw http.ResponseWriter, r *http.Request)`

Handles requests to find a single record in the database.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

### `(a *DBApi) Where(rw http.ResponseWriter, r *http.Request)`

Handles requests to query records based on specific criteria.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

### `(a *DBApi) Update(rw http.ResponseWriter, r *http.Request)`

Handles requests to update a record in the database.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

### `(a *DBApi) AddNewField(rw http.ResponseWriter, r *http.Request)`

Handles requests to add a new field to a record in the database.

- `rw` (http.ResponseWriter): The HTTP response writer.
- `r` (*http.Request): The HTTP request.

