package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Actividad struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Horario     string `json:"horario"`
	Ubicacion   string `json:"ubicacion"`
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./actividades.db")
	if err != nil {
		log.Fatal(err)
	}
	query := `
    CREATE TABLE IF NOT EXISTS actividades (
        id TEXT PRIMARY KEY,
        nombre TEXT,
        descripcion TEXT,
        horario TEXT,
        ubicacion TEXT
    );
    `
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getActividades(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		rows, err := db.Query("SELECT id, nombre, descripcion, horario, ubicacion FROM actividades")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var actividades []Actividad
		for rows.Next() {
			var actividad Actividad
			if err := rows.Scan(&actividad.ID, &actividad.Nombre, &actividad.Descripcion, &actividad.Horario, &actividad.Ubicacion); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			actividades = append(actividades, actividad)
		}
		json.NewEncoder(w).Encode(actividades)
	}
}

func getActividadByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/actividades/")
		if id == "" {
			http.Error(w, "El ID es requerido", http.StatusBadRequest)
			return
		}
		var actividad Actividad
		err := db.QueryRow("SELECT id, nombre, descripcion, horario, ubicacion FROM actividades WHERE id = ?", id).Scan(&actividad.ID, &actividad.Nombre, &actividad.Descripcion, &actividad.Horario, &actividad.Ubicacion)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Actividad no encontrada", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actividad)
	}
}

func createActividad(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var actividad Actividad
		if err := json.NewDecoder(r.Body).Decode(&actividad); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := db.Exec("INSERT INTO actividades (id, nombre, descripcion, horario, ubicacion) VALUES (?, ?, ?, ?, ?)",
			actividad.ID, actividad.Nombre, actividad.Descripcion, actividad.Horario, actividad.Ubicacion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(actividad)
	}
}

func updateActividad(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "El ID es requerido", http.StatusBadRequest)
			return
		}
		var actividad Actividad
		if err := json.NewDecoder(r.Body).Decode(&actividad); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := db.Exec("UPDATE actividades SET nombre = ?, descripcion = ?, horario = ?, ubicacion = ? WHERE id = ?",
			actividad.Nombre, actividad.Descripcion, actividad.Horario, actividad.Ubicacion, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(actividad)
	}
}

func deleteActividad(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "El ID es requerido", http.StatusBadRequest)
			return
		}
		result, err := db.Exec("DELETE FROM actividades WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if rowsAffected == 0 {
			http.Error(w, "Actividad no encontrada", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Modificada la función insertActividades para verificar si la tabla está vacía
func insertActividades(db *sql.DB) {
	// Comprobar si la tabla está vacía
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM actividades").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		return // No insertar si ya hay datos
	}

	actividades := []Actividad{
		{ID: "1", Nombre: "Fútbol", Descripcion: "Entrenamiento de fútbol para principiantes y avanzados.", Horario: "Lunes y Miércoles 16:00 - 18:00", Ubicacion: "Cancha deportiva"},
		{ID: "2", Nombre: "Robótica", Descripcion: "Club de robótica para desarrollar proyectos innovadores.", Horario: "Martes y Jueves 14:00 - 16:00", Ubicacion: "Laboratorio de Ingeniería"},
		{ID: "3", Nombre: "Teatro", Descripcion: "Grupo de teatro donde se practican diversas obras.", Horario: "Viernes 15:00 - 18:00", Ubicacion: "Auditorio Principal"},
		{ID: "4", Nombre: "Baile Folklórico", Descripcion: "Práctica de danzas tradicionales mexicanas.", Horario: "Lunes y Miércoles 17:00 - 19:00", Ubicacion: "Sala de usos múltiples"},
		{ID: "5", Nombre: "Pintura", Descripcion: "Taller de pintura y expresión artística.", Horario: "Martes 10:00 - 12:00", Ubicacion: "Salón de Arte"},
		{ID: "6", Nombre: "Gimnasia", Descripcion: "Clases de gimnasia artística para todos los niveles.", Horario: "Jueves 16:00 - 18:00", Ubicacion: "Gimnasio"},
		{ID: "7", Nombre: "Fotografía", Descripcion: "Taller de fotografía para aprender técnicas básicas y avanzadas.", Horario: "Sábado 10:00 - 13:00", Ubicacion: "Estudio de Arte"},
		{ID: "8", Nombre: "Guitarra", Descripcion: "Clases de guitarra para principiantes y avanzados.", Horario: "Miércoles 10:00 - 12:00", Ubicacion: "Sala de Música"},
		{ID: "9", Nombre: "Cocina", Descripcion: "Taller de cocina internacional.", Horario: "Sábado 15:00 - 18:00", Ubicacion: "Cocina central"},
		{ID: "10", Nombre: "Natación", Descripcion: "Clases de natación para todas las edades.", Horario: "Martes y Jueves 07:00 - 09:00", Ubicacion: "Piscina olímpica"},
		{ID: "11", Nombre: "Ajedrez", Descripcion: "Club de ajedrez para todos los niveles.", Horario: "Jueves 16:00 - 18:00", Ubicacion: "Sala de Juegos"},
		{ID: "12", Nombre: "Ciclismo", Descripcion: "Paseos en bicicleta para principiantes y avanzados.", Horario: "Domingo 07:00 - 10:00", Ubicacion: "Entrada principal"},
		{ID: "13", Nombre: "Dibujo", Descripcion: "Clases de dibujo artístico.", Horario: "Martes 16:00 - 18:00", Ubicacion: "Salón de Arte"},
		{ID: "14", Nombre: "Club de Lectura", Descripcion: "Reuniones para discutir y analizar libros.", Horario: "Viernes 16:00 - 18:00", Ubicacion: "Biblioteca"},
		{ID: "15", Nombre: "Carpintería", Descripcion: "Taller de carpintería para crear proyectos con madera.", Horario: "Sábado 09:00 - 12:00", Ubicacion: "Taller de Carpintería"},
		{ID: "16", Nombre: "Electricidad", Descripcion: "Clases prácticas sobre instalación eléctrica básica.", Horario: "Lunes y Miércoles 14:00 - 16:00", Ubicacion: "Laboratorio de Ingeniería"},
	}
	for _, actividad := range actividades {
		_, err := db.Exec("INSERT OR IGNORE INTO actividades (id, nombre, descripcion, horario, ubicacion) VALUES (?, ?, ?, ?, ?)",
			actividad.ID, actividad.Nombre, actividad.Descripcion, actividad.Horario, actividad.Ubicacion)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	db := initDB()
	defer db.Close()
	insertActividades(db)
	http.HandleFunc("/actividades", getActividades(db))
	http.HandleFunc("/actividades/", getActividadByID(db)) // Cambiado para manejar consultas por ID
	http.HandleFunc("/crear", createActividad(db))
	http.HandleFunc("/actualizar", updateActividad(db))
	http.HandleFunc("/eliminar", deleteActividad(db))
	log.Println("Servidor escuchando en http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
