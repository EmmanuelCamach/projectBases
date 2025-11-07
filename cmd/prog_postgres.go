package main

import (
	"fmt"
	"log"
	"projectBases/pkg/db"
)

func main() {
	fmt.Println("===== Ejecutando módulo PostgreSQL =====")

	conn, err := db.ConnectPostgres()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	if err := db.RunQueries(conn, "postgres"); err != nil {
		log.Fatalf("Error al ejecutar consultas PostgreSQL: %v", err)
	}
}
