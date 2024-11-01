package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Estructuras de datos
type Actividad struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Horario     string `json:"horario"`
	Ubicacion   string `json:"ubicacion"`
}

type Estudiante struct {
	ID               string `json:"id"`
	Nombre           string `json:"nombre"`
	Semestres        string `json:"semestres"`
	Carrera          string `json:"carrera"`
	ActividadID      string `json:"actividad_id"`

}

// Función para inicializar la base de datos
func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/actividades.db")

	if err != nil {
		log.Fatal(err)
	}
	createTable(db)
	return db
}

// Función para crear tablas
func createTable(db *sql.DB) {
	createActividadesTableSQL := `CREATE TABLE IF NOT EXISTS actividades (
		id TEXT PRIMARY KEY,
		nombre TEXT,
		descripcion TEXT,
		horario TEXT,
		ubicacion TEXT
	);`

	createEstudiantesTableSQL := `CREATE TABLE IF NOT EXISTS estudiantes (
		id TEXT PRIMARY KEY,
		nombre TEXT,
		semestres TEXT,
		carrera TEXT,
		actividad_id TEXT,
		FOREIGN KEY (actividad_id) REFERENCES actividades(id)
	);`

	if _, err := db.Exec(createActividadesTableSQL); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(createEstudiantesTableSQL); err != nil {
		log.Fatal(err)
	}
}

// Inserta actividades iniciales en la base de datos
func insertActividades(db *sql.DB) {
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
// Inserta estudiantes iniciales en la base de datos
func insertEstudiantes(db *sql.DB) {
	estudiantes := []Estudiante{
		{ID: "21240143", Nombre: "Juan Pérez", Semestres: "1", Carrera: "Ingeniería en Sistemas", ActividadID: "1"},
		{ID: "21240144", Nombre: "Ana López", Semestres: "1", Carrera: "Ingeniería Industrial", ActividadID: "2"},
		{ID: "21240145", Nombre: "Carlos García", Semestres: "2", Carrera: "Arquitectura", ActividadID: "3"},
		{ID: "21240146", Nombre: "María Fernández", Semestres: "3", Carrera: "Diseño Gráfico", ActividadID: "4"},
		{ID: "21240147", Nombre: "Luis Hernández", Semestres: "1", Carrera: "Ingeniería Electrónica", ActividadID: "5"},
	}

	for _, estudiante := range estudiantes {
		_, err := db.Exec("INSERT OR IGNORE INTO estudiantes (id, nombre, semestres, carrera, actividad_id) VALUES (?, ?, ?, ?, ?)",
			estudiante.ID, estudiante.Nombre, estudiante.Semestres, estudiante.Carrera, estudiante.ActividadID)
		if err != nil {
			log.Printf("Error inserting student %s: %v", estudiante.Nombre, err)
		}
	}
}

// Funciones de la API para agregar estudiantes
func addEstudiante(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var estudiante Estudiante
		if err := json.NewDecoder(r.Body).Decode(&estudiante); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO estudiantes (id, nombre, semestres, carrera, actividad_id) VALUES (?, ?, ?, ?, ?)",
			estudiante.ID, estudiante.Nombre, estudiante.Semestres, estudiante.Carrera, estudiante.ActividadID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(estudiante)
	}
}



// Funciones de la API
func getActividades(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actividades)
	}
}

func getActividadByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		row := db.QueryRow("SELECT id, nombre, descripcion, horario, ubicacion FROM actividades WHERE id = ?", id)

		var actividad Actividad
		if err := row.Scan(&actividad.ID, &actividad.Nombre, &actividad.Descripcion, &actividad.Horario, &actividad.Ubicacion); err != nil {
			http.Error(w, "Actividad no encontrada", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actividad)
	}
}

func addActividad(db *sql.DB) http.HandlerFunc {
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
		var actividad Actividad
		if err := json.NewDecoder(r.Body).Decode(&actividad); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

			_, err := db.Exec("UPDATE actividades SET nombre = ?, descripcion = ?, horario = ?, ubicacion = ? WHERE id = ?",
			actividad.Nombre, actividad.Descripcion, actividad.Horario, actividad.Ubicacion, actividad.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(actividad)
	}
}

func deleteActividad(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM actividades WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func getEstudiantes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT e.id, e.nombre, e.semestres, e.carrera, e.actividad_id FROM estudiantes e LEFT JOIN actividades a ON e.actividad_id = a.id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var estudiantes []Estudiante
		for rows.Next() {
			var estudiante Estudiante
			// Actualiza la línea de escaneo para incluir el nombre de la actividad
			if err := rows.Scan(&estudiante.ID, &estudiante.Nombre, &estudiante.Semestres, &estudiante.Carrera, &estudiante.ActividadID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			estudiantes = append(estudiantes, estudiante)
		}

		w.Header().Set("Content-Type", "application/json") // Asegúrate de establecer el tipo de contenido
		json.NewEncoder(w).Encode(estudiantes) // Envía la respuesta en formato JSON
	}
}




func deleteEstudiante(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_, err := db.Exec("DELETE FROM estudiantes WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func getEstudiantesPorActividadID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actividadID := r.URL.Query().Get("actividad_id")
		rows, err := db.Query("SELECT e.id, e.nombre, e.semestres, e.carrera, e.actividad_id FROM estudiantes e WHERE e.actividad_id = ?", actividadID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var estudiantes []Estudiante
		for rows.Next() {
			var estudiante Estudiante
			if err := rows.Scan(&estudiante.ID, &estudiante.Nombre, &estudiante.Semestres, &estudiante.Carrera, &estudiante.ActividadID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			estudiantes = append(estudiantes, estudiante)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(estudiantes)
	}
}
func updateEstudiante(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var estudiante Estudiante
		if err := json.NewDecoder(r.Body).Decode(&estudiante); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE estudiantes SET nombre = ?, semestres = ?, carrera = ?, actividad_id = ? WHERE id = ?",
			estudiante.Nombre, estudiante.Semestres, estudiante.Carrera, estudiante.ActividadID, estudiante.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(estudiante)
	}
}


func main() {
	db := initDB()
	defer db.Close()

	// Inserta actividades iniciales
	insertActividades(db)
	insertEstudiantes(db)

	http.HandleFunc("/actividades", getActividades(db))
	http.HandleFunc("/actividades/get", getActividadByID(db))
	http.HandleFunc("/actividades/add", addActividad(db))
	http.HandleFunc("/actividades/update", updateActividad(db))
	http.HandleFunc("/actividades/delete", deleteActividad(db))
	http.HandleFunc("/estudiantes", getEstudiantes(db))
	http.HandleFunc("/estudiantes/add", addEstudiante(db))
	http.HandleFunc("/estudiantes/delete", deleteEstudiante(db))
	http.HandleFunc("/estudiantes/actividad", getEstudiantesPorActividadID(db))
	http.HandleFunc("/estudiantes/update", updateEstudiante(db))

	log.Println("Servidor escuchando en http://localhost:7070...")
	if err := http.ListenAndServe("0.0.0.0:7070", nil); err != nil {
		log.Fatal(err)
	}
}
