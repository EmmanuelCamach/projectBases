package main

import (
    "database/sql"
    _ "github.com/sijms/go-ora/v2"
    "fmt"
)

func main() {
    connStr := "oracle://system:1234@localhost:1522/XEPDB1"
    db, err := gosql.Open("oracle", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // DDL
    _, err = db.Exec(`CREATE TABLE users (id NUMBER GENERATED ALWAYS AS IDENTITY, name VARCHAR2(100))`)
    if err != nil {
        panic(err)
    }

    // DML
    _, err = db.Exec(`INSERT INTO users (name) VALUES (:1)`, "Hombre")
    if err != nil {
        panic(err)
    }

    // PL/SQL
    _, err = db.Exec(`
        BEGIN
            DBMS_OUTPUT.PUT_LINE('Ejecutando bloque PL/SQL');
        END;
    `)
    if err != nil {
        panic(err)
    }

    fmt.Println("Conexión y ejecución exitosa en Oracle")
}
