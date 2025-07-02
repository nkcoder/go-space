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