package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockService struct{}

func (s *mockService) GetAllBySeller(sellerID string) ([]Product, error) {
	return []Product{
		{
			ID:          "mock",
			SellerID:    sellerID,
			Description: "generic product",
			Price:       123.55,
		},
	}, nil
}

func TestHandler_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	svc := &mockService{}
	h := NewHandler(svc)
	
	router := gin.Default()
	router.GET("/api/v1/products", h.GetProducts)
	
	t.Run("valid request", func(t *testing.T) {
		req, _ := http.NewRequewst("GET", "/api/v1/products?seller_id=FEX112AC", nil)
		resp := httptest.NewRecorder()
		
		router.ServeHTTP(resp, req)
		
		assert.Equal(t, http.StatusOK, resp.Code)
		
		var products []Product
		json.NewDecoder(resp.Body).Decode(&products)
		
		expected := []Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}
		
		assert.Equal(t, expected, products)
	})
	
	t.Run("missing query parameter", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/products", nil)
		resp := httptest.NewRecorder()
		
		router.ServeHTTP(resp, req)
		
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		
		var body map[string]string
		json.NewDecoder(resp.Body).Decode(&body)
		
		expected := map[string]string{
			"error": "seller_id query param is required",
		}
		
		assert.Equal(t, expected, body)
	})
}
