package handlers

import (
	"net/http"
	"sync"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

// In-memory data store
var (
	users = make(map[string]models.User)
	mu    sync.Mutex // Mutex para proteger el acceso al mapa
)

// Create a new user
func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock() // Bloquear el acceso al mapa
	users[newUser.ID] = newUser
	mu.Unlock() // Desbloquear el acceso al mapa

	c.JSON(http.StatusCreated, newUser)
}

// Get all users
func GetUsers(c *gin.Context) {
	mu.Lock() // Bloquear el acceso al mapa
	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	mu.Unlock() // Desbloquear el acceso al mapa

	c.JSON(http.StatusOK, userList)
}

// Get a single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")

	mu.Lock() // Bloquear el acceso al mapa
	user, exists := users[id]
	mu.Unlock() // Desbloquear el acceso al mapa

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update a user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()         // Bloquear el acceso al mapa
	defer mu.Unlock() // Asegurarse de desbloquear al final

	if _, exists := users[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	updatedUser.ID = id
	users[id] = updatedUser

	c.JSON(http.StatusOK, updatedUser)
}

// Delete a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	mu.Lock()         // Bloquear el acceso al mapa
	defer mu.Unlock() // Asegurarse de desbloquear al final

	if _, exists := users[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	delete(users, id)

	c.JSON(http.StatusNoContent, nil)
}
