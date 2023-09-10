# ToyDB Database

ToyDB is a lightweight JSON database designed to simplify data storage and retrieval for small to medium-scale applications. This README.md file provides an overview of ToyDB's features and usage.

## Features

### JSON-Based
ToyDB uses JSON as the primary data format, making it easy to work with structured data. JSON is human-readable, making it straightforward to inspect and manipulate data.

### File System Storage
Data is stored directly on the file system, eliminating the need for a separate database server. This approach simplifies data management and reduces overhead.

### Collections
ToyDB organizes data into collections, allowing you to group related data together. Each collection acts as a container for JSON documents.

### Query Capabilities
You can perform queries on your data using simple and intuitive commands, enabling efficient data retrieval.


### Configurable
ToyDB is highly configurable, allowing you to specify the location of your database and customize various settings to meet your application's requirements.

## Getting Started

To get started with ToyDB, follow these steps:

1. Clone this repository to your local machine.
2. Install ToyDB by running the installation script or following the installation instructions in the documentation.
3. Create a new database or use an existing one.
4. Use the command-line interface (CLI) to manage and query your data.

For detailed usage instructions, refer to the documentation in the repository.

## Usage

## Node  js Client 
#### `To access Node js client click `

## https://github.com/K-logeshwaran/toydb-query (nodejs client)

### Flags

- `-location`: Specifies the location of your database. Default value is `./database`.
- `-port`: Sets the port for the database API. Default value is `2080`.
- `-serve`: Starts the server on the given port if provided.

### Running the Server

To start the ToyDB server, use the `-serve` flag. For example:

```bash
./toydb -serve
```
By default, the server will listen on http://localhost:2080.

## API Endpoints

`/: Welcome page with information about ToyDB`

`/collection: Endpoint for managing collections.`

`/records: Endpoint for managing records.`

`/findone: Endpoint for finding a specific record.`

`/where: Endpoint for querying records based on criteria.`

`/update: Endpoint for updating records.`

`/addField: Endpoint for adding a new field to a record.`

## Logging
ToyDB logs its activities to a file located at `{dbLoc}/logger.log` file.
## Example
Here's an example of how to start the ToyDB server:
```bash
./your-program -serve -location ./your-database-location -port 8080
```
This will start the server on http://localhost:8080.

Feel free to explore ToyDB and use it to manage your data efficiently.

```

Please replace `./your-program`, `./your-database-location`, and `8080` with the actual values you use when running your program.
```

## Documentation

For more information on how to use ToyDB, consult the official documentation [here](link-to-your-documentation).

## Contributing

We welcome contributions from the community! If you have suggestions, bug reports, or want to contribute code, please open an issue or submit a pull request on this repository.

## License

ToyDB is open-source and released under the [License Name] license. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

We would like to thank the open-source community for their support and contributions to ToyDB.

---

Thank you for choosing ToyDB. Simplify your data management with ease!
