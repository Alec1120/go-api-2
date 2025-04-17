package tests

import (
	"go-api/src/handlers"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"
)

func setupTestDB() *gorm.DB {
	// Create an in-memory SQLite database for testing
	// This will not persist any data and is suitable for unit tests
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// This will create the table in the in-memory database
	db.AutoMigrate(&handlers.Todo{})
	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	// Set Gin to test mode to avoid logging during tests
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	// and set up the routes for the Todo API
	r := gin.Default()
	
	// Set up the routes for the Todo API
	r.GET("/todos", handlers.GetTodos(db))
	r.POST("/todos", handlers.CreateTodo(db))
	r.PUT("/todos/:id", handlers.UpdateTodo(db))
	r.DELETE("/todos/:id", handlers.DeleteTodo(db))
	return r
}

func TestTodoLifecycle(t *testing.T) {
	// Create a new in-memory SQLite database for testing
	db := setupTestDB()
	router := setupRouter(db)

	// Create a POST request to create a new Todo item
	todo := handlers.Todo{Task: "Testify Integration", Done: false}
	body, _ := json.Marshal(todo)

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var created handlers.Todo
	json.Unmarshal(w.Body.Bytes(), &created)
	
	assert.Equal(t, todo.Task, created.Task)
	assert.Equal(t, todo.Done, created.Done)
	assert.NotEmpty(t, created.ID)
	assert.NotEqual(t, 0, created.ID)

	// Create a GET request to retrieve the Todo item
	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var todos []handlers.Todo
	json.Unmarshal(w.Body.Bytes(), &todos)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, created.ID, todos[0].ID)
	assert.Equal(t, created.Task, todos[0].Task)
	assert.Equal(t, created.Done, todos[0].Done)

	// Create a PUT request to update the Todo item
	todos[0].Done = true
	todos[0].Task = "Updated Task"
	updatedBody, _ := json.Marshal(todos[0])
	
	req = httptest.NewRequest(
		http.MethodPut, "/todos/"+strconv.Itoa(todos[0].ID), bytes.NewBuffer(updatedBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, updatedBody, w.Body.Bytes())

	// Create a DELETE request to delete the Todo item
	req = httptest.NewRequest(http.MethodDelete, "/todos/"+strconv.Itoa(todos[0].ID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

