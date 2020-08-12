package controllers

import (
	"github.com/labbsr0x/githunter-api/services"
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
