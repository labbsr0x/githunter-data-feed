package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/labbsr0x/githunter-api/services"

	"github.com/gofiber/fiber"
)

func TestCodeController_GetCommentsHandler_Error_GetComments_Invalid_Body(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	mockIds := []string{}
	mockIds = append(mockIds, "a")

	// Mocking the values Expected
	mockContractService.EXPECT().GetComments(
		mockIds,
		"provider",
		"token",
	).Return(nil, fmt.Errorf("PARSER BODY: error parser input body"))

	app := fiber.New()
	app.Post("/comments", controller.GetCommentsHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	requestJSONStr := `
   {
		"invalid":
			[
				"rafamarts"
			]
		
   }`

	// http.Request
	req := httptest.NewRequest(fiber.MethodPost, "/comments?"+q.Encode(), bytes.NewReader([]byte(requestJSONStr)))
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

func TestCodeController_GetCommentsHandler_Error_GetComments_Invalid_AccessToken(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	mockIds := []string{}
	mockIds = append(mockIds, "a")

	// Mocking the values Expected
	mockContractService.EXPECT().GetComments(
		mockIds,
		"provider",
		"",
	).Return(nil, fmt.Errorf("GetComments invalid token auth code."))

	app := fiber.New()
	app.Post("/comments", controller.GetCommentsHandler)

	q := url.Values{}
	q.Add("access_token", "")
	q.Add("provider", "provider")

	requestJSONStr := `
	   {
			"ids":
				[
					"a"
				]
			
	   }`

	// http.Request
	req := httptest.NewRequest(fiber.MethodPost, "/comments?"+q.Encode(), bytes.NewReader([]byte(requestJSONStr)))
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

func TestCodeController_GetCommentsHandler_Error_GetComments_Unknown_Provider(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	mockIds := []string{}
	mockIds = append(mockIds, "a", "b")

	// Mocking the values Expected
	mockContractService.EXPECT().GetComments(
		mockIds,
		"",
		"token",
	).Return(nil, fmt.Errorf("GetComments unknown provider."))

	app := fiber.New()
	app.Post("/comments", controller.GetCommentsHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "")

	requestJSONStr := `
   {
		"ids":
			[
				"a",
				"b"
			]
		
   }`

	// http.Request
	req := httptest.NewRequest(fiber.MethodPost, "/comments?"+q.Encode(), bytes.NewReader([]byte(requestJSONStr)))
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

func TestCodeController_GetCommentsHandler_Error_GetComments_Invalid_BodyContent(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	mockIds := []string{}
	mockIds = append(mockIds, "a")

	// Mocking the values Expected
	mockContractService.EXPECT().GetComments(
		mockIds,
		"provider",
		"token",
	).Return(nil, fmt.Errorf("GetComments invalid body content."))

	app := fiber.New()
	app.Post("/comments", controller.GetCommentsHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	requestJSONStr := `
   {
		"ids":
			[
				"a"
			]
		
   }`

	// http.Request
	req := httptest.NewRequest(fiber.MethodPost, "/comments?"+q.Encode(), bytes.NewReader([]byte(requestJSONStr)))
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

func TestCodeController_GetCommentsHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr :=
		`{
			"data": [{
				"name": "name",
				"owner": "owner",
				"createdAt": "1995-10-10T00:00:01Z",
				"number": 1,
				"url": "https://github.com/test/comments/url",
				"id": "MDEyOklzc3VlQ29tbWVudDcxMjIyNzc4NA==",
				"author": "author"
			}]
		}`
	mockResponse := &services.CommentsResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	mockIds := []string{}
	mockIds = append(mockIds, "a")

	// Mocking the values Expected
	mockContractService.EXPECT().GetComments(
		mockIds,
		"provider",
		"token",
	).Return(mockResponse, nil)

	app := fiber.New()
	app.Post("/comments", controller.GetCommentsHandler)

	q := url.Values{}
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	requestJSONStr := `
   {
		"ids":
			[
				"a"
			]
		
   }`

	// http.Request
	req := httptest.NewRequest(fiber.MethodPost, "/comments?"+q.Encode(), bytes.NewReader([]byte(requestJSONStr)))
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

	theContract := &services.CommentsResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
