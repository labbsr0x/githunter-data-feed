package controllers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labbsr0x/githunter-data-feed/services/mock"
)

func GetMockContractServiceAndController(t *testing.T) (m *mock.MockContract, c *Controller) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	return mockContractService, controller
}
