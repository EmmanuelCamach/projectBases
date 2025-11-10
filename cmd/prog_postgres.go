package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=1117496331 dbname=projectBases sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// DDL
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT)`)
	if err != nil {
		panic(err)
	}

	// DML
	_, err = db.Exec(`INSERT INTO users (name) VALUES ($1)`, "Hombre")
	if err != nil {
		panic(err)
	}

	// PL/pgSQL
	_, err = db.Exec(`
        DO $$
        BEGIN
            RAISE NOTICE 'Ejecutando bloque PL/pgSQL';
        END $$;
    `)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión y ejecución exitosa en PostgreSQL")
}
