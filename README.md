## Conexión e interacción cdesde Go con Oracle y PostgreSQL
Este proyecto conecta una aplicación escrita en Go con dos motores de base de datos: Oracle y PostgreSQL, permitiendo realizar operaciones (consultas, inserciones, actualizaciones, etc.) sobre ambas desde un mismo entorno.


## Estructura
- /sql: scripts de creación e inserción
- /pkg/db: conexión y ejecución
- /cmd: programas principales
- /tests: validaciones

## Ejecución
go run ./cmd/prog_postgres.go
go run ./cmd/prog_oracle.go
