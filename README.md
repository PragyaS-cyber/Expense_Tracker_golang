This is a simple golang project which describes how API works and it has used golang's "GoFr" framework.


*HTTP APIs*
An HTTP API is an API that uses Hypertext Transfer Protocol as the communication protocol between the two systems. HTTP APIs expose endpoints as API gateways for HTTP requests to have access to a server.
example: HTTP API every time you set a Zoom meeting in your Google calendar. The API defines how Zoom can communicate directly with Google’s servers to embed a Zoom meeting into the event rather than having to copy and paste the meeting invitation into a field.

*REST APIs*
Rest api stands for representational state transfer and is architechtural pattern for creating web services.REST is a ruleset that defines best practices for sharing data between clients and the server. It’s essentially a design style used when creating HTTP or other APIs that asks you to use CRUD functions only, regardless of the complexity. 
 REST emphasizes the scalability of components and the simplicity of interfaces.

*CRUD operations:*
Create
Read
Update
Delete

other APIS:
GraphQL
Falcor
gRPC


package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/vikash/gofr/pkg/gofr"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Task model
type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Database setup
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Task{})
}

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", nil)
}

// CRUD operations

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updatedTask Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	db.Model(&task).Updates(updatedTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	db.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
