package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// Configuración de conexión
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1117496331"
	dbname   = "projectBases"
)

func main() {
	fmt.Println("=== PROYECTO GO - POSTGRESQL ===")
	fmt.Println("Iniciando conexión con PostgreSQL...")

	// Construir cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abrir conexión
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir conexión:", err)
	}
	defer db.Close()

	// Verificar conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	fmt.Println("✓ Conexión exitosa a PostgreSQL")
	time.Sleep(1 * time.Second)

	// 1. LIMPIAR SCHEMA ANTERIOR SI EXISTE
	fmt.Println("--- PASO 1: Limpiando schema anterior ---")
	cleanupSchema(db)
	time.Sleep(1 * time.Second)

	// 2. CREAR TABLAS (DDL)
	fmt.Println("\n--- PASO 2: Creando tablas (DDL) ---")
	createTables(db)
	time.Sleep(1 * time.Second)

	// 3. LISTAR TABLAS
	fmt.Println("\n--- PASO 3: Listando tablas creadas ---")
	listTables(db)
	time.Sleep(1 * time.Second)

	// 4. DESCRIBIR TABLA
	fmt.Println("\n--- PASO 4: Describiendo estructura de tabla 'estudiantes' ---")
	describeTable(db, "estudiantes")
	time.Sleep(1 * time.Second)

	// 5. INSERTAR DATOS (DML)
	fmt.Println("\n--- PASO 5: Insertando datos (DML - INSERT) ---")
	insertData(db)
	time.Sleep(1 * time.Second)

	// 6. CONSULTAR DATOS
	fmt.Println("\n--- PASO 6: Consultando datos (DML - SELECT) ---")
	selectData(db)
	time.Sleep(1 * time.Second)

	// 7. ACTUALIZAR DATOS
	fmt.Println("\n--- PASO 7: Actualizando datos (DML - UPDATE) ---")
	updateData(db)
	time.Sleep(1 * time.Second)

	// 8. ELIMINAR DATOS
	fmt.Println("\n--- PASO 8: Eliminando datos (DML - DELETE) ---")
	deleteData(db)
	time.Sleep(1 * time.Second)

	// 9. CREAR Y LLAMAR FUNCIÓN
	fmt.Println("\n--- PASO 9: Creando y llamando función PL/pgSQL ---")
	createAndCallFunction(db)
	time.Sleep(1 * time.Second)

	// 10. CREAR Y LLAMAR PROCEDIMIENTO
	fmt.Println("\n--- PASO 10: Creando y llamando procedimiento PL/pgSQL ---")
	createAndCallProcedure(db)
	time.Sleep(1 * time.Second)

	// 11. CREAR Y PROBAR TRIGGER
	fmt.Println("\n--- PASO 11: Creando y probando trigger (disparador) ---")
	createAndTestTrigger(db)
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== PROYECTO FINALIZADO EXITOSAMENTE ===")
}

func cleanupSchema(db *sql.DB) {
	queries := []string{
		"DROP TABLE IF EXISTS auditoria CASCADE",
		"DROP TABLE IF EXISTS inscripciones CASCADE",
		"DROP TABLE IF EXISTS cursos CASCADE",
		"DROP TABLE IF EXISTS estudiantes CASCADE",
		"DROP TABLE IF EXISTS departamentos CASCADE",
		"DROP FUNCTION IF EXISTS calcular_promedio_curso(INTEGER)",
		"DROP PROCEDURE IF EXISTS actualizar_cupos(INTEGER, INTEGER)",
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Advertencia al limpiar: %v\n", err)
		}
	}
	fmt.Println("✓ Schema anterior limpiado")
}

func createTables(db *sql.DB) {
	// DDL para crear 5 tablas
	ddlStatements := []string{
		`CREATE TABLE departamentos (
			id_departamento SERIAL PRIMARY KEY,
			nombre VARCHAR(100) NOT NULL,
			edificio VARCHAR(50)
		)`,
		`CREATE TABLE estudiantes (
			id_estudiante SERIAL PRIMARY KEY,
			nombre VARCHAR(100) NOT NULL,
			apellido VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE,
			fecha_ingreso DATE,
			id_departamento INTEGER REFERENCES departamentos(id_departamento)
		)`,
		`CREATE TABLE cursos (
			id_curso SERIAL PRIMARY KEY,
			nombre_curso VARCHAR(100) NOT NULL,
			creditos INTEGER,
			cupo_maximo INTEGER,
			id_departamento INTEGER REFERENCES departamentos(id_departamento)
		)`,
		`CREATE TABLE inscripciones (
			id_inscripcion SERIAL PRIMARY KEY,
			id_estudiante INTEGER REFERENCES estudiantes(id_estudiante),
			id_curso INTEGER REFERENCES cursos(id_curso),
			calificacion DECIMAL(3,2),
			fecha_inscripcion DATE
		)`,
		`CREATE TABLE auditoria (
			id_auditoria SERIAL PRIMARY KEY,
			tabla VARCHAR(50),
			operacion VARCHAR(10),
			fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			descripcion TEXT
		)`,
	}

	for i, ddl := range ddlStatements {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Printf("Error al crear tabla %d: %v", i+1, err)
		} else {
			fmt.Printf("✓ Tabla %d creada exitosamente\n", i+1)
		}
	}
}

func listTables(db *sql.DB) {
	query := `SELECT table_name FROM information_schema.tables 
			  WHERE table_schema = 'public' ORDER BY table_name`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar tablas: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("Tablas en el schema:")
	count := 0
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		count++
		fmt.Printf("  %d. %s\n", count, tableName)
	}
}

func describeTable(db *sql.DB, tableName string) {
	query := `SELECT column_name, data_type, is_nullable 
			  FROM information_schema.columns 
			  WHERE table_name = $1 ORDER BY ordinal_position`

	rows, err := db.Query(query, tableName)
	if err != nil {
		log.Printf("Error al describir tabla: %v", err)
		return
	}
	defer rows.Close()

	fmt.Printf("Estructura de la tabla '%s':\n", tableName)
	fmt.Println("  Columna           | Tipo         | Null")
	fmt.Println("  ------------------|--------------|------")

	for rows.Next() {
		var colName, dataType, nullable string
		rows.Scan(&colName, &dataType, &nullable)
		fmt.Printf("  %-17s | %-12s | %s\n", colName, dataType, nullable)
	}
}

func insertData(db *sql.DB) {
	// Insertar departamentos
	depts := []string{
		"INSERT INTO departamentos (nombre, edificio) VALUES ('Ingeniería', 'Edificio A')",
		"INSERT INTO departamentos (nombre, edificio) VALUES ('Ciencias', 'Edificio B')",
		"INSERT INTO departamentos (nombre, edificio) VALUES ('Humanidades', 'Edificio C')",
	}

	for _, stmt := range depts {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("Error en INSERT: %v", err)
		}
	}
	fmt.Println("✓ Departamentos insertados")

	// Insertar estudiantes
	students := []string{
		"INSERT INTO estudiantes (nombre, apellido, email, fecha_ingreso, id_departamento) VALUES ('Juan', 'Pérez', 'juan@email.com', '2023-01-15', 1)",
		"INSERT INTO estudiantes (nombre, apellido, email, fecha_ingreso, id_departamento) VALUES ('María', 'González', 'maria@email.com', '2023-02-20', 1)",
		"INSERT INTO estudiantes (nombre, apellido, email, fecha_ingreso, id_departamento) VALUES ('Carlos', 'Rodríguez', 'carlos@email.com', '2023-03-10', 2)",
		"INSERT INTO estudiantes (nombre, apellido, email, fecha_ingreso, id_departamento) VALUES ('Ana', 'Martínez', 'ana@email.com', '2023-04-05', 2)",
		"INSERT INTO estudiantes (nombre, apellido, email, fecha_ingreso, id_departamento) VALUES ('Luis', 'López', 'luis@email.com', '2023-05-12', 3)",
	}

	for _, stmt := range students {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("Error en INSERT: %v", err)
		}
	}
	fmt.Println("✓ Estudiantes insertados")

	// Insertar cursos
	courses := []string{
		"INSERT INTO cursos (nombre_curso, creditos, cupo_maximo, id_departamento) VALUES ('Bases de Datos', 4, 30, 1)",
		"INSERT INTO cursos (nombre_curso, creditos, cupo_maximo, id_departamento) VALUES ('Algoritmos', 4, 25, 1)",
		"INSERT INTO cursos (nombre_curso, creditos, cupo_maximo, id_departamento) VALUES ('Física I', 3, 40, 2)",
	}

	for _, stmt := range courses {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("Error en INSERT: %v", err)
		}
	}
	fmt.Println("✓ Cursos insertados")

	// Insertar inscripciones
	enrollments := []string{
		"INSERT INTO inscripciones (id_estudiante, id_curso, calificacion, fecha_inscripcion) VALUES (1, 1, 4.5, '2023-06-01')",
		"INSERT INTO inscripciones (id_estudiante, id_curso, calificacion, fecha_inscripcion) VALUES (2, 1, 3.8, '2023-06-01')",
		"INSERT INTO inscripciones (id_estudiante, id_curso, calificacion, fecha_inscripcion) VALUES (3, 2, 4.2, '2023-06-02')",
		"INSERT INTO inscripciones (id_estudiante, id_curso, calificacion, fecha_inscripcion) VALUES (1, 2, 4.0, '2023-06-02')",
	}

	for _, stmt := range enrollments {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("Error en INSERT: %v", err)
		}
	}
	fmt.Println("✓ Inscripciones insertadas")
}

func selectData(db *sql.DB) {
	query := `SELECT e.nombre, e.apellido, c.nombre_curso, i.calificacion 
			  FROM estudiantes e
			  JOIN inscripciones i ON e.id_estudiante = i.id_estudiante
			  JOIN cursos c ON i.id_curso = c.id_curso
			  ORDER BY e.apellido`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error en SELECT: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("Estudiantes y sus inscripciones:")
	fmt.Println("  Nombre           | Curso              | Calificación")
	fmt.Println("  -----------------|--------------------|--------------")

	for rows.Next() {
		var nombre, apellido, curso string
		var calificacion float64
		rows.Scan(&nombre, &apellido, &curso, &calificacion)
		fullName := nombre + " " + apellido
		fmt.Printf("  %-17s | %-18s | %.2f\n", fullName, curso, calificacion)
	}
}

func updateData(db *sql.DB) {
	// Actualizar calificación
	query := "UPDATE inscripciones SET calificacion = 4.7 WHERE id_inscripcion = 1"
	result, err := db.Exec(query)
	if err != nil {
		log.Printf("Error en UPDATE: %v", err)
		return
	}

	rows, _ := result.RowsAffected()
	fmt.Printf("✓ Actualización completada: %d registro(s) modificado(s)\n", rows)

	// Mostrar el registro actualizado
	var nombre, curso string
	var calif float64
	query = `SELECT e.nombre || ' ' || e.apellido, c.nombre_curso, i.calificacion
			 FROM inscripciones i
			 JOIN estudiantes e ON i.id_estudiante = e.id_estudiante
			 JOIN cursos c ON i.id_curso = c.id_curso
			 WHERE i.id_inscripcion = 1`

	err = db.QueryRow(query).Scan(&nombre, &curso, &calif)
	if err == nil {
		fmt.Printf("  Nuevo valor: %s en %s = %.2f\n", nombre, curso, calif)
	}
}

func deleteData(db *sql.DB) {
	// Contar antes de borrar
	var countBefore int
	db.QueryRow("SELECT COUNT(*) FROM inscripciones").Scan(&countBefore)
	fmt.Printf("Inscripciones antes del DELETE: %d\n", countBefore)

	// Eliminar inscripciones con calificación menor a 4.0
	query := "DELETE FROM inscripciones WHERE calificacion < 4.0"
	result, err := db.Exec(query)
	if err != nil {
		log.Printf("Error en DELETE: %v", err)
		return
	}

	rows, _ := result.RowsAffected()
	fmt.Printf("✓ Eliminación completada: %d registro(s) eliminado(s)\n", rows)

	// Contar después de borrar
	var countAfter int
	db.QueryRow("SELECT COUNT(*) FROM inscripciones").Scan(&countAfter)
	fmt.Printf("Inscripciones después del DELETE: %d\n", countAfter)
}

func createAndCallFunction(db *sql.DB) {
	// Crear función PL/pgSQL
	functionDDL := `
		CREATE OR REPLACE FUNCTION calcular_promedio_curso(p_id_curso INTEGER)
		RETURNS DECIMAL(3,2) AS $$
		DECLARE
			v_promedio DECIMAL(3,2);
		BEGIN
			SELECT AVG(calificacion) INTO v_promedio
			FROM inscripciones
			WHERE id_curso = p_id_curso;
			
			RETURN COALESCE(v_promedio, 0.0);
		END;
		$$ LANGUAGE plpgsql;
	`

	_, err := db.Exec(functionDDL)
	if err != nil {
		log.Printf("Error al crear función: %v", err)
		return
	}
	fmt.Println("✓ Función 'calcular_promedio_curso' creada")

	// Llamar la función
	var promedio float64
	query := "SELECT calcular_promedio_curso(1)"
	err = db.QueryRow(query).Scan(&promedio)
	if err != nil {
		log.Printf("Error al llamar función: %v", err)
		return
	}

	fmt.Printf("✓ Función ejecutada - Promedio del curso 1: %.2f\n", promedio)
}

func createAndCallProcedure(db *sql.DB) {
	// Crear procedimiento PL/pgSQL
	procedureDDL := `
		CREATE OR REPLACE PROCEDURE actualizar_cupos(
			p_id_curso INTEGER,
			p_nuevo_cupo INTEGER
		) AS $$
		BEGIN
			UPDATE cursos 
			SET cupo_maximo = p_nuevo_cupo 
			WHERE id_curso = p_id_curso;
			
			INSERT INTO auditoria (tabla, operacion, descripcion)
			VALUES ('cursos', 'UPDATE', 'Actualización de cupo para curso ' || p_id_curso);
		END;
		$$ LANGUAGE plpgsql;
	`

	_, err := db.Exec(procedureDDL)
	if err != nil {
		log.Printf("Error al crear procedimiento: %v", err)
		return
	}
	fmt.Println("✓ Procedimiento 'actualizar_cupos' creado")

	// Mostrar cupo antes
	var cupoAntes int
	db.QueryRow("SELECT cupo_maximo FROM cursos WHERE id_curso = 1").Scan(&cupoAntes)
	fmt.Printf("  Cupo antes: %d\n", cupoAntes)

	// Llamar el procedimiento
	_, err = db.Exec("CALL actualizar_cupos(1, 35)")
	if err != nil {
		log.Printf("Error al llamar procedimiento: %v", err)
		return
	}
	fmt.Println("✓ Procedimiento ejecutado")

	// Mostrar cupo después
	var cupoDespues int
	db.QueryRow("SELECT cupo_maximo FROM cursos WHERE id_curso = 1").Scan(&cupoDespues)
	fmt.Printf("  Cupo después: %d\n", cupoDespues)

	// Verificar registro en auditoría
	var desc string
	db.QueryRow("SELECT descripcion FROM auditoria ORDER BY id_auditoria DESC LIMIT 1").Scan(&desc)
	fmt.Printf("  Registro de auditoría: %s\n", desc)
}

func createAndTestTrigger(db *sql.DB) {
	// Crear función para el trigger
	triggerFunctionDDL := `
		CREATE OR REPLACE FUNCTION fn_auditoria_estudiantes()
		RETURNS TRIGGER AS $$
		BEGIN
			IF TG_OP = 'INSERT' THEN
				INSERT INTO auditoria (tabla, operacion, descripcion)
				VALUES ('estudiantes', 'INSERT', 'Nuevo estudiante: ' || NEW.nombre || ' ' || NEW.apellido);
			END IF;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`

	_, err := db.Exec(triggerFunctionDDL)
	if err != nil {
		log.Printf("Error al crear función de trigger: %v", err)
		return
	}
	fmt.Println("✓ Función de trigger creada")

	// Crear trigger
	triggerDDL := `
		CREATE TRIGGER trg_auditoria_estudiantes
		AFTER INSERT ON estudiantes
		FOR EACH ROW
		EXECUTE FUNCTION fn_auditoria_estudiantes();
	`

	_, err = db.Exec(triggerDDL)
	if err != nil {
		log.Printf("Error al crear trigger: %v", err)
		return
	}
	fmt.Println("✓ Trigger 'trg_auditoria_estudiantes' creado")

	// Probar el trigger insertando un estudiante
	insertQuery := `INSERT INTO estudiantes (nombre, apellido, email, fecha_ingreso, id_departamento) 
					VALUES ('Pedro', 'Sánchez', 'pedro@email.com', CURRENT_DATE, 1)`

	_, err = db.Exec(insertQuery)
	if err != nil {
		log.Printf("Error al insertar estudiante: %v", err)
		return
	}
	fmt.Println("✓ Nuevo estudiante insertado (trigger activado)")

	// Verificar que el trigger funcionó
	var descripcion string
	query := "SELECT descripcion FROM auditoria WHERE tabla = 'estudiantes' ORDER BY id_auditoria DESC LIMIT 1"
	err = db.QueryRow(query).Scan(&descripcion)
	if err == nil {
		fmt.Printf("  Registro automático en auditoría: %s\n", descripcion)
	}
}
