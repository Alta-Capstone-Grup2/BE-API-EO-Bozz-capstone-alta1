package service

import (
	"capstone-alta1/features/city"
	"capstone-alta1/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.CityRepository)
	returnData := []city.Core{{ID: 1, CityName: "ada"}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].CityName, response[0].CityName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get All", func(t *testing.T) {
		// var tx *gorm.DB
		repo.On("GetAll").Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}
