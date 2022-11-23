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
// @Router /liveness [get]
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
// @Router /liveness [get]
func HealthController(c *gin.Context) {
	c.JSON(200, HealthResponse{
		Status: "OK",
	})
}
