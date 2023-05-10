// Completar los métodos del handler para que pasen las pruebas. (desde return func etc)

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
		var req request
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		h.shark.Configure([2]float64{req.XPosition, req.YPosition}, req.Speed)
		context.JSON(200, response{Success: true})
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
		var req request
		err := context.ShouldBindJSON(&req)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		h.prey.SetSpeed(req.Speed)
		context.JSON(200, response{Success: true})
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
		err, timeToCatch := h.shark.Hunt(h.prey)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		context.JSON(200, response{Success: true, Time: timeToCatch})
	}
}
