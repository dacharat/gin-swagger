package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

// NewUserHandler create new user handler
func NewUserHandler() UserHandler {
	return UserHandler{}
}

// CreateUserHandler create user handler
func (u UserHandler) CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetUserHandler get user handler
func (u UserHandler) GetUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateUserHandler update user handler
func (u UserHandler) UpdateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUserHandler delete user handler
func (u UserHandler) DeleteUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
