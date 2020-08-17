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

func TestReposController_GetReposHandler_Error_GetRepos_Invalid_AccessToken(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetLastRepos(
		"",
		"provider",
	).Return(nil, fmt.Errorf("GetLastRepos invalid auth token."))

	app := fiber.New()
	app.Get("/repos", controller.GetReposHandler)

	q := url.Values{}
	q.Add("access_token", "")
	q.Add("provider", "provider")

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

func TestReposController_GetReposHandler_Error_GetRepos_Unknown_Provider(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetLastRepos(
		"token",
		"",
	).Return(nil, fmt.Errorf("GetLastRepos unknown provider."))

	app := fiber.New()
	app.Get("/repos", controller.GetReposHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "")

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

func TestCodeController_GetReposHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr :=
		`{
			"name": "fakeNameUser",
			"repositories": [
				"fakeRepo01",
				"fakeRepo02",
				"fakeRepo03",
				"fakeRepo04",
				"fakeRepo05",
				"fakeRepo06",
				"fakeRepo07",
				"fakeRepo08",
				"fakeRepo09",
				"fakeRepo10",
			]
	}`
	mockResponse := &services.ReposResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetLastRepos(
		"token",
		"provider",
	).Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/repos", controller.GetReposHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/repos?"+q.Encode(), nil)
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

	theContract := &services.ReposResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
