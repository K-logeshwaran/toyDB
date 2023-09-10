# Driver Package

The `Driver` package provides essential functionalities for ToyDB, a lightweight JSON database solution.

## Types
- DataBase
- Collection
- Wrapper


# `DataBase Struct`

`DataBase` is a struct representing the main database instance in ToyDB.

- `Location`: The location of the database.
- `collections`: An instance of the `Collection` interface.

### `NewDB(loc string, logger string, col Collection) *DataBase`

Creates a new `DataBase` instance.

- `loc`: The location of the database.
- `logger`: The location of the log file.
- `col`: An instance of the `Collection` interface.

## Methods

### `CreateCollection(name string) error`

Creates a new collection in the database.

- `name`: The name of the collection to create.

### `IsCollectionExist(name string) bool`

Checks if a collection with the given name exists.

- `name`: The name of the collection to check.

### `IsCollectionNotExist(name string) bool`

Checks if a collection with the given name does not exist.

- `name`: The name of the collection to check.

### `PopulateRecords(collection string, data []byte) (message string, err error)`

Populates records in a collection.

- `collection`: The name of the collection to populate.
- `data`: The data to populate the collection with.

### `createuuid() string`

Generates a new UUID string.

### `ReadAll(collection string, limit int, don chan bool, resultCh chan Wrapper) (chan Wrapper, error)`

Reads all records from a collection with a specified limit.

- `collection`: The name of the collection to read from.
- `limit`: The maximum number of records to read.
- `don`: A channel to indicate when the operation is done.
- `resultCh`: A channel to receive the retrieved records.

### `ReadAllGPt(collection string, limit int, done chan bool, resultCh chan Wrapper) error`

Reads all records from a collection using Goroutines with a specified limit.

- `collection`: The name of the collection to read from.
- `limit`: The maximum number of records to read.
- `done`: A channel to indicate when the operation is done.
- `resultCh`: A channel to receive the retrieved records.

### `FindOneById(collection string, id string) (*Wrapper, string, error)`

Finds a record in a collection by its ID.

- `collection`: The name of the collection to search.
- `id`: The ID of the record to find.

### `UpdateOneById(collection, id, field string, value interface{}) (*Wrapper, error)`

Updates a record in a collection by its ID.

- `collection`: The name of the collection to update.
- `id`: The ID of the record to update.
- `field`: The field to update.
- `value`: The new value for the field.

### `AddField(collection, id, field string, value interface{}) (*Wrapper, error)`

Adds a new field with a value to a record in a collection.

- `collection`: The name of the collection to update.
- `id`: The ID of the record to update.
- `field`: The new field to add.
- `value`: The value to set for the new field.

### `commit(recordpath string, w *Wrapper)`

Commits changes to a record.

- `recordpath`: The path to the record file.
- `w`: The record to commit.

### `ListCollections() *Wrapper`

Lists all collections in the database.

### `Where(collection string, field string, value interface{}) ([]Wrapper, error)`

Queries records in a collection where a specific field matches a value.

- `collection`: The name of the collection to query.
- `field`: The field to match.
- `value`: The value to match against the specified field.

# `Collection Struct`

Represents a collection of strings.

- **Collections** (Field): An array of strings representing the collection names.

## Functions

### `CreateCollectionFiles(loc string)`

Creates a collections JSON file at the specified location if it does not already exist.

- `loc` (string): The directory path where the collections JSON file should be created.

### `NewCollection(dbloc string) Collection`

Creates a new `Collection` instance by reading data from a collections JSON file.

- `dbloc` (string): The directory path where the collections JSON file is located.

- Returns: A `Collection` instance populated with data from the JSON file.

### `(c *Collection) AddCollection(cl string)`

Adds a new collection name to the `Collection` instance.

- `c` (*Collection): The `Collection` instance to which the new collection name should be added.
- `cl` (string): The name of the collection to be added.

### `(c *Collection) Commit(dbloc string)`

Commits changes made to the `Collection` instance by writing data back to the collections JSON file.

- `c` (*Collection): The `Collection` instance to be committed.
- `dbloc` (string): The directory path where the collections JSON file should be saved.


# `Wrapper Struct`

The `Wrapper` type is a map of string keys to interface values, providing flexibility in constructing and manipulating JSON-like structures.

## Functions

### `BuildWrapper(Data []byte) *Wrapper`

Creates and returns a new `Wrapper` instance by unmarshaling JSON data.

- `Data` ([]byte): The JSON data to be used for building the `Wrapper` structure.
- Returns: A pointer to a `Wrapper` instance.

### `(w *Wrapper) AddField(name string, value interface{}) *Wrapper`

Adds a new field with a specified name and value to the `Wrapper`.

- `w` (*Wrapper): The `Wrapper` instance to which the field should be added.
- `name` (string): The name of the field.
- `value` (interface{}): The value of the field.
- Returns: A pointer to the modified `Wrapper` instance.

### `(w *Wrapper) Update(name string, value interface{}) *Wrapper`

Updates an existing field in the `Wrapper` with a new value.

- `w` (*Wrapper): The `Wrapper` instance to be updated.
- `name` (string): The name of the field to be updated.
- `value` (interface{}): The new value for the field.
- Returns: A pointer to the modified `Wrapper` instance.

### `(w *Wrapper) Value() map[string]interface{}`

Returns the underlying map of the `Wrapper`.

- Returns: A map[string]interface{} containing the data of the `Wrapper`.

### `(w *Wrapper) ToBytes() []byte`

Serializes the `Wrapper` to a JSON byte array.

- Returns: A byte slice containing the JSON representation of the `Wrapper`.

### `(w *Wrapper) ToJson() string`

Serializes the `Wrapper` to a JSON string.

- Returns: A string containing the JSON representation of the `Wrapper`.

### `WrapperArrayToBytes(W []Wrapper) []byte`

Serializes an array of `Wrapper` instances to a JSON byte array.

- `W` ([]Wrapper): An array of `Wrapper` instances to be serialized.
- Returns: A byte slice containing the JSON representation of the array of `Wrapper` instances.



