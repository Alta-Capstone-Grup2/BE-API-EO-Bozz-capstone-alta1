package service

import (
	"capstone-alta1/features/discussion"
	"capstone-alta1/mocks"
	"errors"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// var err error

func TestCreate(t *testing.T) {
	repo := new(mocks.DiscussionRepository)
	inputData := discussion.Core{Comment: "ada", ClientID: 1, PartnerID: 1, ServiceID: 1}
	var c echo.Context
	t.Run("Success Create", func(t *testing.T) {
		repo.On("Create", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputData, c)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	// t.Run("Failed to Create", func(t *testing.T) {
	// 	inputData := discussion.Core{Comment: "ada", ClientID: 1, PartnerID: 1, ServiceID: 1}
	// 	repo.On("Create").Return(errors.New("error")).Once()
	// 	srv := New(repo)
	// 	err := srv.Create(inputData, c)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, "error", err.Error())
	// 	repo.AssertExpectations(t)
	// })
}
func TestGetAll(t *testing.T) {
	repo := new(mocks.DiscussionRepository)
	returnData := []discussion.Core{{ID: 1, Comment: "ada", PartnerID: 1, ClientID: 1, ServiceID: 1, CreatedAt: time.Now()}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Comment, response[0].Comment)
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

func TestGetById(t *testing.T) {
	repo := new(mocks.DiscussionRepository)
	returnData := discussion.Core{ID: 1, Comment: "ada", PartnerID: 1, ClientID: 1, ServiceID: 1, CreatedAt: time.Now()}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetById", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Comment, response.Comment)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetById", id).Return(discussion.Core{}, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}
