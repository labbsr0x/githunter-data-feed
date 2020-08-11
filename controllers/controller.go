package controllers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labbsr0x/githunter-api/services"
	"github.com/labbsr0x/githunter-api/services/mock"
)

type Controller struct {
	Contract services.Contract
}

func NewController() *Controller {
	theController := &Controller{
		Contract: services.New(),
	}
	return theController
}

func GetMockContractServiceAndController(t *testing.T) (m *mock.MockContract, c *Controller) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	return mockContractService, controller
}
