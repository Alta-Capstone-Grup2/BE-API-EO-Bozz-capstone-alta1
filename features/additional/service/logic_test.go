package service

import (
	"capstone-alta1/features/additional"
	"capstone-alta1/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.AdditionalRepository)
	inputData := additional.Core{AdditionalName: "ada", AdditionalPrice: 1, PartnerID: 1}
	t.Run("Success Create", func(t *testing.T) {
		repo.On("Create", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	// t.Run("Failed Create", func(t *testing.T) {
	// 	repo.On("Create", inputData).Return(errors.New("failed")).Once()
	// 	srv := New(repo)
	// 	err := srv.Create(inputData)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, "failed", err.Error())
	// 	repo.AssertExpectations(t)
	// })
}
func TestGetAll(t *testing.T) {
	repo := new(mocks.AdditionalRepository)
	returnData := []additional.Core{{ID: 1, AdditionalName: "ada", AdditionalPrice: 1, PartnerID: 1}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].AdditionalName, response[0].AdditionalName)
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

func TestUpdate(t *testing.T) {
	repo := new(mocks.AdditionalRepository)
	inputData := additional.Core{AdditionalName: "ada", AdditionalPrice: 1, PartnerID: 1}
	var id uint = 1
	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", inputData, id).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(inputData, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed Update", func(t *testing.T) {
		repo.On("Update", inputData, id).Return(errors.New("error")).Once()
		srv := New(repo)
		err := srv.Update(inputData, id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.AdditionalRepository)
	var id uint = 1
	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", id).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete", func(t *testing.T) {
		repo.On("Delete", id).Return(errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
