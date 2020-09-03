package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/labbsr0x/githunter-api/services"
)

func TestUserController_GetUserHandler_Error_GetUserStats_Invalid_Login(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserStats(
		"",
		"token",
		"provider",
	).Return(nil, fmt.Errorf("GetUserStats invalid name of user."))

	app := fiber.New()
	app.Get("/user/stats", controller.GetUserHandler)

	q := url.Values{}
	q.Add("login", "")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/user/stats?"+q.Encode(), nil)
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

func TestUserController_GetUserHandler_Error_GetUserStats_Invalid_AccessToken(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserStats(
		"login",
		"",
		"provider",
	).Return(nil, fmt.Errorf("GetUserStats invalid token auth code."))

	app := fiber.New()
	app.Get("/user/stats", controller.GetUserHandler)

	q := url.Values{}
	q.Add("login", "login")
	q.Add("access_token", "")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/user/stats?"+q.Encode(), nil)
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

func TestUserController_GetUserHandler_Error_GetUserStats_Unknown_Provider(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserStats(
		"login",
		"token",
		"",
	).Return(nil, fmt.Errorf("GetUserStats unknown provider."))

	app := fiber.New()
	app.Get("/user/stats", controller.GetUserHandler)

	q := url.Values{}
	q.Add("login", "login")
	q.Add("access_token", "token")
	q.Add("provider", "")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/user/stats?"+q.Encode(), nil)
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

func TestUserController_GetUserHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr :=
		`{
			"name": "Name",
			"login": "login",
			"amount": {
					"repositories": 666,
					"commits": 666,
					"pullRequests": 666,
					"issues": 666,
					"starsReceived": 666,
					"followers": 666
			}
	}`
	mockResponse := &services.UserResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserStats(
		"login",
		"token",
		"provider",
	).Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/user/stats", controller.GetUserHandler)

	q := url.Values{}
	q.Add("login", "login")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/user/stats?"+q.Encode(), nil)
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

	theContract := &services.UserResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
