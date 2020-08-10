package controllers

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/golang/mock/gomock"
	"github.com/labbsr0x/githunter-api/services/mock"
)

func TestCodeController_GetCodeHandler_Error_GetCode_Unknown_Provider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	reposController := &CodeController{
		Contract: mockContractService,
	}

	// Mocking the values Expected
	mockContractService.EXPECT().GetLastRepos(, "token", "providerTest").Return(nil, fmt.Errorf("GetLastRepos unknown provider: providerTest"))

	app := fiber.New()
	app.Get("/repos", reposController.GetReposHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "providerTest")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/repos?"+q.Encode(), nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// http.Response
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("expect response status code %d got %d", fiber.StatusInternalServerError, resp.StatusCode)
	}

}
