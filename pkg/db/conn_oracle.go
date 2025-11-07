package db

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

// ConnectOracle crea y devuelve una conexión activa a Oracle.
func ConnectOracle() (*sql.DB, error) {
	connStr := "user=estudiante password=sistemas connectString=localhost:1521/UsuarioEstudiante"

	db, err := sql.Open("godror", connStr)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión a Oracle: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error conectando a Oracle: %v", err)
	}

	fmt.Println("✅ Conectado correctamente a Oracle")
	return db, nil
}
