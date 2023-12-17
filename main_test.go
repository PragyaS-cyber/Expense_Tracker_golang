package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
    "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	clearTable()

	payload := []byte(`{"title":"Test Task", "description":"Test Description", "completed":false}`)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	assert.Equal(t, http.StatusCreated, response.Code)

	var task Task
	json.NewDecoder(response.Body).Decode(&task)

	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "Test Description", task.Description)
	assert.Equal(t, false, task.Completed)
}

func TestGetTasks(t *testing.T) {
	clearTable()
	addTasks(5)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code)

	var tasks []Task
	json.NewDecoder(response.Body).Decode(&tasks)

	assert.Equal(t, 5, len(tasks))
}

// Add more test functions for other CRUD operations...

// Helper functions

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", createTaskCmd).Methods("POST")
	router.HandleFunc("/tasks", getTasksCmd).Methods("GET")
	router.ServeHTTP(rr, req)
	return rr
}

func clearTable() {
	db.Delete(&Task{})
}

func addTasks(count int) {
	for i := 0; i < count; i++ {
		db.Create(&Task{Title: "Task " + string(i), Description: "Description " + string(i), Completed: false})
	}
}
