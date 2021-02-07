package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dacharat/gin-swagger/cmd/requests"
	"github.com/dacharat/gin-swagger/cmd/responses"
)

// UserHandler user handler
type UserHandler struct{}

// NewUserHandler create new user handler
func NewUserHandler() UserHandler {
	return UserHandler{}
}

// CreateUserHandler create user handler
// @summary Create user
// @description create new user
// @tags users
// @security ApiKeyAuth
// @id CreateUserHandler
// @accept json
// @produce json
// @Param X-ID header string true "With x-id header"
// @param User body requests.CreateUserRequest true "Data for create user"
// @response 200 {object} responses.CreateUserResponse "OK"
// @response 400 {object} map[string]string "Bad Request"
// @router /user [post]
func (u UserHandler) CreateUserHandler(c *gin.Context) {
	id := c.GetString("ID")
	var req requests.CreateUserRequest

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
	}

	c.JSON(http.StatusOK, responses.CreateUserResponse{
		Username: fmt.Sprintf("%s-%s", id, req.Username),
	})
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
