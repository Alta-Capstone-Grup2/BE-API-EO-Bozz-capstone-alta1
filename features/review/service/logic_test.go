package service

import (
	"capstone-alta1/features/review"
	"capstone-alta1/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.ReviewRepository)
	t.Run("Success Create", func(t *testing.T) {
		inputRepo := review.Core{Review: "ada", Rating: 2, OrderID: 1, ServiceID: 1}
		repo.On("Create", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(inputRepo)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

}
