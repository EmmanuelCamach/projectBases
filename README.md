## Proyecto Bases de Datos
Este proyecto implementa dos programas en Go que se conectan a Oracle y PostgreSQL respectivamente, ejecutan un conjunto de consultas predefinidas, y verifican la equivalencia de resultados entre ambas bases.


## Estructura
- /sql: scripts de creación e inserción
- /pkg/db: conexión y ejecución
- /cmd: programas principales
- /tests: validaciones

## Ejecución
go run ./cmd/prog_postgres.go
go run ./cmd/prog_oracle.go
