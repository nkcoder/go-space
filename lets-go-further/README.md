> next: Advaned CRUD Operations => Managing SQL Query Timeouts

## Notes

- When there is a valid `go.mod` file in the root of your project directory, your project is a module.

## Project structure

```
- bin
- cmd
  - api
    - main.go
- internal
- migrations
- remote
- go.mod
- Makefile
```

- bin: contains compiled application binaries, ready for deployment to a production server
- cmd/api: contains the application-specific code for the API, including the code running the server, reading and
  writing HTTP requests, and authentication
- internal: contains ancillary packages, including database, data validation, send email etc. Any code which isn't
  application-specific and can potentially be reused will live here. Special meaning: any packages which lives under
  this directory can only be imported by code inside the parent of this directory
- migration: SQL migration files
- go.mod: declares project dependencies, versions and module path

## Error handling

- **Expected errors**: Errors that can occur during normal operation, such as a database query timeout, unavailable
  network resources, or invalid user input. These are typically caused by factors outside your program's control. It is
  best practice to return and handle these errors gracefully.
- **Unexpected errors**: Errors that should not happen during normal operation, often due to developer mistakes or logic
  errors in the codebase. These are exceptional cases where using `panic` is more acceptable. The Go standard library
  often panics in such situations, for example, when accessing an out-of-bounds slice index or closing an already-closed
  channel.

## Format & Run

```sh
auto/format
auto/run
```

## Usage

Create a movie:
```sh
curl -d '{"title": "123", "runtime": "107 mins", "year": 201, "genres": []}' localhost:4000/v1/movies
{
  "movie": {
    "title": "123",
    "year": 201,
    "runtime": 107,
    "genres": []
  }
}
```

Get a movie:
```shell
curl localhost:4000/v1/movies/2
{
  "movie": {
    "id": 2,
    "title": "Casablanca",
    "runtime": 180,
    "genres": [
      "drama",
      "romance",
      "war"
    ],
    "version": 1
  }
}
```

## Database

postgres://greenlight:Password@2025@localhost:54322/greenlight
