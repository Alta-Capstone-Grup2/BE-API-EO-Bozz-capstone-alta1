package service

import (
	"capstone-alta1/features/review"
	"capstone-alta1/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.ReviewRepository)
	inputData := review.Core{Review: "ada", Rating: 5, ClientID: 1, OrderID: 1, ServiceID: 1}
	t.Run("Success Create", func(t *testing.T) {
		repo.On("Create", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	// t.Run("Failed to Create", func(t *testing.T) {
	// 	inputData := review.Core{Review: "ada", Rating: 2, ClientID: 1, OrderID: 1, ServiceID: 1}
	// 	repo.On("Create", inputData).Return(errors.New("error")).Once()
	// 	srv := New(repo)
	// 	err := srv.Create(inputData)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, "error", err.Error())
	// 	repo.AssertExpectations(t)
	// })
}
func TestGetAll(t *testing.T) {
	repo := new(mocks.ReviewRepository)
	returnData := []review.Core{{ID: 1, Review: "ada", Rating: 5, ClientID: 1, OrderID: 1, ServiceID: 1}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Review, response[0].Review)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get All", func(t *testing.T) {
		// var tx *gorm.DB
		repo.On("GetAll").Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.ReviewRepository)
	returnData := review.Core{ID: 1, Review: "ada", Rating: 5, ClientID: 1, OrderID: 1, ServiceID: 1}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetById", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Review, response.Review)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetById", id).Return(review.Core{}, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}
