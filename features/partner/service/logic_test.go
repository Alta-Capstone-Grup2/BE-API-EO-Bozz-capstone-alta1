package service

import (
	"capstone-alta1/features/partner"
	"capstone-alta1/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	returnData := []partner.Core{{ID: 1, CompanyName: "ada", CreatedAt: time.Now(), User: partner.UserCore{Name: "ada"}, VerificationStatus: "ada", UserID: 1}}
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", "query").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].CompanyName, response[0].CompanyName)
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
	repo := new(mocks.PartnerRepository)
	returnData := partner.Core{ID: 1, User: partner.UserCore{Name: "ada", Email: "ada", Role: "ada"}, PICPosition: "ada", PICPhone: "ada", PICAddress: "ada", CompanyName: "ada", CompanyPhone: "ada", CompanyCity: "ada", CompanyImageFile: "ada", CompanyAddress: "ada", LinkWebsite: "ada", NIBNumber: "ada", NIBImageFile: "ada", SIUPNumber: "ada", SIUPImageFile: "ada", Event1Name: "ada", Event1ImageFile: "ada", Event2Name: "ada", Event2ImageFile: "ada", Event3Name: "ada", Event3ImageFile: "ada", BankName: "ada", BankAccountNumber: "ada", BankAccountName: "ada", VerificationStatus: "ada", VerificationLog: "ada", UserID: 1}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetById", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.PICPosition, response.PICPosition)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetById", id).Return(partner.Core{}, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetById(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.PartnerRepository)
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

func TestGetServiceById(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	returnData := []partner.ServiceCore{{ID: 1, ServiceName: "ada", ServiceDescription: "ada", ServiceIncluded: "ada", ServiceCategory: "ada", ServicePrice: 1, AverageRating: 1, ServiceImageFile: "ada", City: "ada", PartnerID: 1}}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetServices", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetServices(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].ServiceName, response[0].ServiceName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetServices", id).Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetServices(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestGetOrderById(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	returnData := []partner.OrderCore{{ID: 1, EventName: "ada", StartDate: time.Now(), EndDate: time.Now(), EventLocation: "ada", EventAddress: "ada", NoteForPartner: "ada", ServiceName: "ada", ServicePrice: 1, GrossAmmount: 1, PaymentMethod: "ada", OrderStatus: "ada", PayoutRecieptFile: "ada", PayoutDate: time.Now(), ServiceID: 1, ClientID: 1}}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetOrders", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetOrders(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].ServiceName, response[0].ServiceName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetOrders", id).Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetOrders(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestGetAdditionalById(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	returnData := []partner.AdditionalCore{{ID: 1, AdditionalName: "ada", AdditionalPrice: 1, PartnerID: 1}}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetAdditionals", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAdditionals(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].AdditionalName, response[0].AdditionalName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetAdditionals", id).Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetAdditionals(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestOrderConfirm(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success confirm order", func(t *testing.T) {
		repo.On("UpdateOrderConfirmStatus", id, id).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateOrderConfirmStatus(id, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed confirm  order", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("UpdateOrderConfirmStatus", id, id).Return(errors.New("error")).Once()
		srv := New(repo)
		err := srv.UpdateOrderConfirmStatus(id, id)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetPartnerRegister(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	returnData := []partner.Core{{ID: 1, CompanyName: "ada", CreatedAt: time.Now(), User: partner.UserCore{Name: "ada"}, VerificationStatus: "ada", UserID: 1}}
	// var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get partner register", func(t *testing.T) {
		repo.On("GetPartnerRegisterData", "ada", "ada", "ada").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetPartnerRegisterData("ada", "ada", "ada")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].CompanyName, response[0].CompanyName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get partner regis", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetPartnerRegisterData", "ada", "ada", "ada").Return(nil, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetPartnerRegisterData("ada", "ada", "ada")
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}

func TestGetByIdRegister(t *testing.T) {
	repo := new(mocks.PartnerRepository)
	returnData := partner.Core{ID: 1, User: partner.UserCore{Name: "ada", Email: "ada", Role: "ada"}, PICPosition: "ada", PICPhone: "ada", PICAddress: "ada", CompanyName: "ada", CompanyPhone: "ada", CompanyCity: "ada", CompanyImageFile: "ada", CompanyAddress: "ada", LinkWebsite: "ada", NIBNumber: "ada", NIBImageFile: "ada", SIUPNumber: "ada", SIUPImageFile: "ada", Event1Name: "ada", Event1ImageFile: "ada", Event2Name: "ada", Event2ImageFile: "ada", Event3Name: "ada", Event3ImageFile: "ada", BankName: "ada", BankAccountNumber: "ada", BankAccountName: "ada", VerificationStatus: "ada", VerificationLog: "ada", UserID: 1}
	var id uint = 1
	// var id2 uint = 0
	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("GetPartnerRegisterDataByID", id).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetPartnerRegisterDataByID(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.PICPosition, response.PICPosition)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get By Id", func(t *testing.T) {
		// returnDatabyid := discussion.Core{}
		repo.On("GetPartnerRegisterDataByID", id).Return(partner.Core{}, errors.New("error")).Once()
		srv := New(repo)
		response, err := srv.GetPartnerRegisterDataByID(id)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData, response)
		repo.AssertExpectations(t)
	})
}
