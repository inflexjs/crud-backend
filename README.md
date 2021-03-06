# crud-backend


## Run web server:
```
go run .\cmd\main.go
```

## Guidebook

- `cmd/` - starting point of the project. 


- `configs/` - non-classified project data.  


- `internal/` - application core.
    - `handler/` - application transport layer.  
    - `service/` - application business logic.  
    - `storage/` - working with the database.


- `models/` - description of structures and types.


- `pkg/` - independent packages.


- `schema/` - database migrate schema.


- `.env.example` - example for using secure config.


- `server.go` - default net/http server cfg and methods.
