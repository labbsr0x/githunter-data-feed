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
	"github.com/labbsr0x/githunter-data-feed/services"
)

func TestCodeController_GetCommitsHandler_Error_GetCommitsRepo_Invalid_NameAndOwner(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetCommitsRepo(
		"",
		"",
		"token",
		"provider",
	).Return(nil, fmt.Errorf("GetCommitsRepo invalid path of repository."))

	app := fiber.New()
	app.Get("/commits", controller.GetCommitsHandler)

	q := url.Values{}
	q.Add("name", "")
	q.Add("owner", "")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/commits?"+q.Encode(), nil)
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

func TestCodeController_GetCommitsHandler_Error_GetCommitsRepo_Invalid_AccessToken(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetCommitsRepo(
		"name",
		"owner",
		"",
		"provider",
	).Return(nil, fmt.Errorf("GetCommitsRepo invalid token auth code."))

	app := fiber.New()
	app.Get("/commits", controller.GetCommitsHandler)

	q := url.Values{}
	q.Add("name", "name")
	q.Add("owner", "owner")
	q.Add("access_token", "")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/commits?"+q.Encode(), nil)
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

func TestCodeController_GetCommitsHandler_Error_GetCommitsRepo_Unknown_Provider(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetCommitsRepo(
		"name",
		"owner",
		"token",
		"",
	).Return(nil, fmt.Errorf("GetCommitsRepo unknown provider."))

	app := fiber.New()
	app.Get("/commits", controller.GetCommitsHandler)

	q := url.Values{}
	q.Add("name", "name")
	q.Add("owner", "owner")
	q.Add("access_token", "token")
	q.Add("provider", "")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/commits?"+q.Encode(), nil)
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

func TestCodeController_GetCommitsHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr :=
		`{
			"commits": [
				{
					"message": "fakeMessageCommit",
					"committedDate": "1995-10-10T00:00:01Z"
				}
			]
		}`
	mockResponse := &services.CommitsResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetCommitsRepo(
		"name",
		"owner",
		"token",
		"provider",
	).Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/commits", controller.GetCommitsHandler)

	q := url.Values{}
	q.Add("name", "name")
	q.Add("owner", "owner")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/commits?"+q.Encode(), nil)
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

	theContract := &services.CommitsResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
