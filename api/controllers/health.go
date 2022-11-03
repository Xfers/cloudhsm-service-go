package controllers

// Health check controller

import (
	"github.com/gin-gonic/gin"
)

// HealthResponse godoc
// @Summary Health check response
// @Description health check response
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
type HealthResponse struct {
	Status string `json:"status"`
}

// Health godoc
// @Summary Health check
// @Description health check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
