package service

import (
	"capstone-alta1/features/client"
	"capstone-alta1/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.ClientRepository)
	returnData := []client.Core{{ID: 1, Gender: "ada", Address: "ada", City: "ada", Phone: "ada", ClientImageFile: "ada", UserID: 1, User: client.User{Name: "ada", Email: "ada", Role: "Client"}}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", "query").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Gender, response[0].Gender)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get All", func(t *testing.T) {
		repo.On("GetAll", "query").Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.ClientRepository)
	returnData := client.Core{ID: 1, Gender: "ada", Address: "ada", City: "ada", Phone: "ada", ClientImageFile: "ada", UserID: 1, User: client.User{Name: "ada", Email: "ada", Role: "Client"}}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetById", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Gender, response.Gender)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetById", id).Return(client.Core{}, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ClientRepository)
	var id uint = 1
	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", id, id).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(id, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete", func(t *testing.T) {
		repo.On("Delete", id, id).Return(errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(id, id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

// func TestUpdate(t *testing.T) {
// 	repo := new(mocks.ClientRepository)
// 	inputData := client.Core{Gender: "ada", Address: "ada", City: "ada", Phone: "ada", ClientImageFile: "ada", UserID: 1, User: client.User{Name: "ada", Email: "ada", Role: "Client", Password: "ada"}}
// 	var id uint = 1
// 	var c echo.Context
// 	// var id2 uint = 0
// 	t.Run("Success Update", func(t *testing.T) {
// 		repo.On("Update", inputData, id, id, c).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Update(inputData, id, id, c)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

// t.Run("failed Updatecomplete order", func(t *testing.T) {
// 	repo.On("UpdateCompleteOrder", inputData, id, id).Return(errors.New("error")).Once()
// 	srv := New(repo)
// 	err := srv.UpdateCompleteOrder(inputData, id, id)
// 	assert.NotNil(t, err)
// 	repo.AssertExpectations(t)
// })
// }

func TestGetOrderById(t *testing.T) {
	repo := new(mocks.ClientRepository)
	returnData := []client.Order{{ID: 1, EventName: "ada", StartDate: time.Now(), EndDate: time.Now(), EventLocation: "ada", ServiceName: "ada", GrossAmmount: 1, OrderStatus: "ada", ServiceID: 1, ClientID: 1}}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetOrderById", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetOrderById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].EventName, response[0].EventName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetOrderById", id).Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetOrderById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestUpdateComplete(t *testing.T) {
	repo := new(mocks.ClientRepository)
	inputData := client.Order{OrderStatus: "ada"}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Updatecomplete order", func(t *testing.T) {
		repo.On("UpdateCompleteOrder", inputData, id, id).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateCompleteOrder(inputData, id, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed Updatecomplete order", func(t *testing.T) {
		repo.On("UpdateCompleteOrder", inputData, id, id).Return(errors.New("error")).Once()
		srv := New(repo)
		err := srv.UpdateCompleteOrder(inputData, id, id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
