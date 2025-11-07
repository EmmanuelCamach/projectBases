package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  //"fmt"
  "log"
)

func main() {
  connStr := "host=localhost port=5432 dbname=mi_db user=mi_user password=mi_pass sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil { log.Fatal(err) }
  defer db.Close()

  // Leer consultas de archivo o embed
  // Ejecutar Q1: SELECT -> obtener filas -> imprimir en formato CSV
  // Ejecutar Q3: UPDATE -> imprimir rowsAffected
  // Ejecutar Q4: DDL -> imprimir success
  // etc.
}
