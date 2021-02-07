package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler health handler
// @summary Health Check
// @description Health checking for the service
// @id HealthHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, time.Now())
}
