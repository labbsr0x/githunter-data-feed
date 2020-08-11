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

func TestCodeController_GetCodeHandler_Error_GetInfoCodePage_Invalid_NameAndOwner(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetInfoCodePage(
		"",
		"",
		"token",
		"provider",
	).Return(nil, fmt.Errorf("GetInfoCodePage invalid path of repository."))

	app := fiber.New()
	app.Get("/code", controller.GetCodeHandler)

	q := url.Values{}
	q.Add("name", "")
	q.Add("owner", "")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/code?"+q.Encode(), nil)
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

func TestCodeController_GetCodeHandler_Error_GetInfoCodePage_Invalid_AccessToken(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetInfoCodePage(
		"name",
		"owner",
		"",
		"provider",
	).Return(nil, fmt.Errorf("GetInfoCodePage wihtout token auth code."))

	app := fiber.New()
	app.Get("/code", controller.GetCodeHandler)

	q := url.Values{}
	q.Add("name", "name")
	q.Add("owner", "owner")
	q.Add("access_token", "")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/code?"+q.Encode(), nil)
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

func TestCodeController_GetCodeHandler_Error_GetInfoCodePage_Unknown_Provider(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetInfoCodePage(
		"name",
		"owner",
		"token",
		"",
	).Return(nil, fmt.Errorf("GetInfoCodePage unknown provider."))

	app := fiber.New()
	app.Get("/code", controller.GetCodeHandler)

	q := url.Values{}
	q.Add("name", "name")
	q.Add("owner", "owner")
	q.Add("access_token", "token")
	q.Add("provider", "")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/code?"+q.Encode(), nil)
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

func TestCodeController_GetCodeHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr :=
		`{
			"name": "fakeNameRepo",
			"description": "fakeDescription",
			"createdAt": "1995-10-10T00:00:01Z",
			"primaryLanguage": "fakeLanguage",
			"repositoryTopics": [
				"fakeTopic"
			],
			"watchers": 999,
			"stars": 999,
			"forks": 999,
			"lastCommitDate": "1995-10-10T00:00:01Z",
			"commits": 999,
			"readme": "fakeTextReadme",
			"contributing": "fakeTextContributing",
			"licenseInfo": "fakeLicenseInfo",
			"codeOfConduct": {
				"body": "fakeBodyCodeOfConduct",
				"resourcePath": "fakeResourcePath"
			},
			"releases": 999,
			"contributors": 999,
			"languages": {
				"quantity": 1,
				"languages": [
					{
						"size": 100,
						"name": "C ANSI"
					}
				]
			},
			"diskUsage": 100
	}`
	mockResponse := &services.CodeResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetInfoCodePage(
		"name",
		"owner",
		"token",
		"provider",
	).Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/code", controller.GetCodeHandler)

	q := url.Values{}
	q.Add("name", "name")
	q.Add("owner", "owner")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/code?"+q.Encode(), nil)
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

	theContract := &services.CodeResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
