package main

import (
	"encoding/json"
	"fmt"
	//"net/http"
	"github.com/jinzhu/gorm"
	"github.com/vikash/gofr/pkg/gofr"
	_ //"github.com/jinzhu/gorm/dialects/sqlite"
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
	// Update connection details based on Dockerized PostgreSQL instance
	db, err = gorm.Open("postgres", "host=postgres user=your_username dbname=your_database sslmode=disable password=your_password")
	if err != nil {
	   panic("Failed to connect to database")
	}
	db.AutoMigrate(&Task{})
}

func main() {
	app := gofr.NewCMD()

	// Subcommands
	app.SubCommand("create", createTaskCmd)
	app.SubCommand("get", getTasksCmd)
	app.SubCommand("get/{id}", getTaskCmd)
	app.SubCommand("update/{id}", updateTaskCmd)
	app.SubCommand("delete/{id}", deleteTaskCmd)

	app.Run()
}

// Subcommand handlers

func createTaskCmd(c *gofr.Context) (interface{}, error) {
	defer c.Request.Body.Close()

	var task Task
	if err := json.NewDecoder(c.Request.Body).Decode(&task); err != nil {
		return nil, err
	}

	db.Create(&task)
	return task, nil
}

func getTasksCmd(c *gofr.Context) (interface{}, error) {
	var tasks []Task
	db.Find(&tasks)
	return tasks, nil
}

func getTaskCmd(c *gofr.Context) (interface{}, error) {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, fmt.Errorf("Task not found")
	}
	return task, nil
}

func updateTaskCmd(c *gofr.Context) (interface{}, error) {
	defer c.Request.Body.Close()

	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, fmt.Errorf("Task not found")
	}

	var updatedTask Task
	if err := json.NewDecoder(c.Request.Body).Decode(&updatedTask); err != nil {
		return nil, err
	}

	db.Model(&task).Updates(updatedTask)
	return task, nil
}

func deleteTaskCmd(c *gofr.Context) (interface{}, error) {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, fmt.Errorf("Task not found")
	}

	db.Delete(&task)
	return nil, nil
}