package main

import (
	"fmt"
	"log"
	"projectBases/pkg/db"
)

func main() {
	fmt.Println("===== Ejecutando módulo Oracle =====")

	conn, err := db.ConnectOracle()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	if err := db.RunQueries(conn, "oracle"); err != nil {
		log.Fatalf("Error al ejecutar consultas Oracle: %v", err)
	}
}
