package products

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}

func TestGetAllBySeller_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := NewService(repo)
	
	expected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	
	repo.On("GetAllBySeller", "FEX112AC").Return(expected, nil)
	
	result, err := svc.GetAllBySeller("FEX112AC")
	
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetAllBySeller_Error(t *testing.T) {
	repo := new(MockRepository)
	svc := NewService(repo)
	
	expectedErr := errors.New("some error")
	
	repo.On("GetAllBySeller", "invalid_id").Return(nil, expectedErr)
	
	result, err := svc.GetAllBySeller("invalid_id")
	
	assert.Error(t, err)
	assert.Nil(t, result)
}
