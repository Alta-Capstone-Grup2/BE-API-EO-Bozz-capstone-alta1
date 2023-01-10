package service

import (
	"capstone-alta1/features/user"
	"capstone-alta1/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCreate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	inputData := user.Core{Name: "ada", Email: "ada", Password: "ada", Role: "Admin"}
// 	var c echo.Context
// 	t.Run("Success Create", func(t *testing.T) {
// 		repo.On("Create", inputData, c).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData, c)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

//		// t.Run("Failed Create", func(t *testing.T) {
//		// 	repo.On("Create", inputData).Return(errors.New("failed")).Once()
//		// 	srv := New(repo)
//		// 	err := srv.Create(inputData)
//		// 	assert.NotNil(t, err)
//		// 	assert.Equal(t, "failed", err.Error())
//		// 	repo.AssertExpectations(t)
//		// })
//	}
func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := []user.Core{{ID: 1, Name: "ada", Email: "ada", Password: "ada", Role: "ada"}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get All", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := user.Core{ID: 1, Name: "ada", Email: "ada", Password: "ada", Role: "ada"}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetById", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetById", id).Return(user.Core{}, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

// func TestUpdate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	inputData := user.Core{Name: "ada", Email: "ada", Password: "ada"}
// 	var id uint = 1
// 	t.Run("Success Update", func(t *testing.T) {
// 		repo.On("Update", inputData, id).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Update(inputData, id)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	// t.Run("failed Update", func(t *testing.T) {
// 	// 	repo.On("Update", inputData, id).Return(errors.New("error")).Once()
// 	// 	srv := New(repo)
// 	// 	err := srv.Update(inputData, id)
// 	// 	assert.NotNil(t, err)
// 	// 	repo.AssertExpectations(t)
// 	// })
// }

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepository)
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

// func TestUpdatePassword(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	inputData := user.Core{Password: "ada"}
// 	var id uint = 1
// 	t.Run("Success Update", func(t *testing.T) {
// 		repo.On("UpdatePassword", inputData, id).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.UpdatePassword(inputData, id)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

// t.Run("failed Update", func(t *testing.T) {
// 	repo.On("Update", inputData, id).Return(errors.New("error")).Once()
// 	srv := New(repo)
// 	err := srv.Update(inputData, id)
// 	assert.NotNil(t, err)
// 	repo.AssertExpectations(t)
// })
// }
