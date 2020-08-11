package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golang/mock/gomock"
	"github.com/labbsr0x/githunter-api/services"
	"github.com/labbsr0x/githunter-api/services/mock"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestController_GetPullsHandler_Error_GetPulls_Unknown_Provider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	// Mocking the values Expected
	mockContractService.EXPECT().GetPulls(10, "", "", "providerTest", "token").Return(nil, fmt.Errorf("'Get' using unknown provider: providerTest"))

	app := fiber.New()
	app.Get("/pulls", controller.GetPullsHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "providerTest")
	q.Add("owner", "")
	q.Add("name", "")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/pulls?"+q.Encode(), nil)
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

func TestController_GetPullsHandler_Error_GetPulls_Invalid_Token(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	// Mocking the values Expected
	mockContractService.EXPECT().GetPulls(10, "", "", "github", "token").Return(nil, fmt.Errorf("'Get' using unknown token: token"))

	app := fiber.New()
	app.Get("/pulls", controller.GetPullsHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "github")
	q.Add("owner", "")
	q.Add("name", "")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/pulls?"+q.Encode(), nil)
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

func TestController_GetPullsHandler_Error_GetPulls_Invalid_NameAndOwner(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	theValidToken := "fakeValidToken"

	// Mocking the values Expected
	mockContractService.EXPECT().GetPulls(10, "fakeOwner", "fakeName", "github", theValidToken).Return(nil, nil)

	app := fiber.New()
	app.Get("/pulls", controller.GetPullsHandler)

	q := url.Values{}
	q.Add("access_token", theValidToken)
	q.Add("provider", "github")
	q.Add("owner", "fakeOwner")
	q.Add("name", "fakeName")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/pulls?"+q.Encode(), nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// http.Response
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusNoContent {
		t.Errorf("expect response status code %d got %d", fiber.StatusNoContent, resp.StatusCode)
	}

}

func TestController_GetPullsHandler_Success(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	theValidToken := "fakeValidToken"
	responseJSONStr := `
    {
        "total": 10,
		"pulls": [
			{
				"number": 1234
			}
		]
    }`
	mockResponse := &services.PullsResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetPulls(10, "fakeOwner", "fakeName", "github", theValidToken).Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/pulls", controller.GetPullsHandler)

	q := url.Values{}
	q.Add("access_token", theValidToken)
	q.Add("provider", "github")
	q.Add("owner", "fakeOwner")
	q.Add("name", "fakeName")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/pulls?"+q.Encode(), nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// http.Response
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expect response status code %d got %d", fiber.StatusOK, resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	theContract := &services.PullsResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
