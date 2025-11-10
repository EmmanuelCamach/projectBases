package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// LoadQueries carga las consultas SQL específicas para el motor indicado.
func LoadQueries(dbType string) ([]string, error) {
	var filePath string

	switch strings.ToLower(dbType) {
	case "postgres":
		filePath = "./sql/consultas_postgres.sql"
	case "oracle":
		filePath = "./sql/consultas_oracle.sql"
	default:
		return nil, fmt.Errorf("tipo de base de datos no reconocido: %s", dbType)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo de consultas: %v", err)
	}

	// Separa las consultas por ';'
	rawQueries := strings.Split(string(data), ";")
	var queries []string
	for _, q := range rawQueries {
		q = strings.TrimSpace(q)
		if q != "" {
			queries = append(queries, q)
		}
	}
	_, err := db.ExecDDL(query string)

	return queries, nil
}

// RunQueries ejecuta todas las consultas cargadas en la base de datos indicada.
func RunQueries(db *sql.DB, dbType string) error {
	fmt.Printf("▶ Ejecutando consultas en %s...\n", strings.Title(dbType))

	queries, err := LoadQueries(dbType)
	if err != nil {
		return err
	}

	for i, query := range queries {
		fmt.Printf("\n[%s] Consulta #%d:\n%s\n", strings.ToUpper(dbType), i+1, query)
		rows, err := db.Query(query)
		if err != nil {
			fmt.Printf("⚠ Error ejecutando consulta %d: %v\n", i+1, err)
			continue
		}
		defer rows.Close()

		cols, _ := rows.Columns()
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range cols {
			valuePtrs[i] = &values[i]
		}

		for rows.Next() {
			err = rows.Scan(valuePtrs...)
			if err != nil {
				fmt.Printf("⚠ Error leyendo fila: %v\n", err)
				continue
			}

			for i, col := range cols {
				fmt.Printf("%s = %v\t", col, values[i])
			}
			fmt.Println()
		}
	}

	fmt.Printf("\n✅ Finalizado correctamente en %s.\n", strings.Title(dbType))
	return nil
}
