package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// ConnectPostgres crea y devuelve una conexión activa a PostgreSQL.
func ConnectPostgres() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=tu_password dbname=mi_basedatos sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión a PostgreSQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error conectando a PostgreSQL: %v", err)
	}

	fmt.Println("✅ Conectado correctamente a PostgreSQL")
	return db, nil
}
