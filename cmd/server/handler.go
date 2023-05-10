package server

import (
	"functional/prey"
	"functional/shark"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
	}
}
