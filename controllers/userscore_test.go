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

func TestUserScoreController_GetUserScoreHandler_Error_GetUserScore_Invalid_Login(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserScore(
		"",
		"token",
		"provider",
	).Return(nil, fmt.Errorf("GetUserScore invalid name of user."))

	app := fiber.New()
	app.Get("/userscore", controller.GetUserScoreHandler)

	q := url.Values{}
	q.Add("login", "")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/userscore?"+q.Encode(), nil)
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

func TestUserScoreController_GetUserScoreHandler_Error_GetUserScore_Invalid_AccessToken(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserScore(
		"login",
		"",
		"provider",
	).Return(nil, fmt.Errorf("GetUserScore invalid token auth code."))

	app := fiber.New()
	app.Get("/userscore", controller.GetUserScoreHandler)

	q := url.Values{}
	q.Add("login", "login")
	q.Add("access_token", "")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/userscore?"+q.Encode(), nil)
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

func TestUserScoreController_GetUserScoreHandler_Error_GetUserScore_Unknown_Provider(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserScore(
		"login",
		"token",
		"",
	).Return(nil, fmt.Errorf("GetUserScore unknown provider."))

	app := fiber.New()
	app.Get("/userscore", controller.GetUserScoreHandler)

	q := url.Values{}
	q.Add("login", "login")
	q.Add("access_token", "token")
	q.Add("provider", "")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/userscore?"+q.Encode(), nil)
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

func TestUserScoreController_GetUserScoreHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr :=
		`{
			"name": "Name",
			"login": "login",
			"id": "id",
			"followers": [
				"tovarlds",
			],
			"organizations": [
				"ibm"
			],
			"contributedRepositories": [
				{
					"name": "opensource.guide",
					"owner": "github"
				},
			],
			"pulls": [
				{
					"number": 39,
					"name": "githunter-web",
					"owner": "labbsr0x",
					"state": "MERGED",
					"createdAt": "2020-12-07T22:32:30Z",
					"updatedAt": "2020-12-09T15:47:54Z",
					"closedAt": "2020-12-09T15:47:50Z",
					"merged": true,
					"mergedAt": "2020-12-09T15:47:50Z",
					"author": "finhaa",
					"labels": null,
					"comments": {
						"totalCount": 0,
						"updatedAt": "",
						"data": null
					},
					"participants": {
						"totalCount": 1,
						"users": [
							"finhaa"
						]
					}
				},
			],
			"issues": [
				{
					"number": 12791,
					"state": "CLOSED",
					"createdAt": "2020-08-06T00:45:59Z",
					"updatedAt": "2020-08-06T11:02:47Z",
					"closedAt": "2020-08-06T11:02:47Z",
					"author": "fakeAuthor",
					"labels": [
						"fakeLabel"
					],
					"totalParticipants": 2,
					"timelineItems": {
						"totalCount": 0,
						"updatedAt": "2020-08-06T11:02:47Z",
						"nodes": null
					}
				},
			],
			"ownedRepositories": [
				{
					"name": "github-slideshow",
					"owner": "finhaa",
					"createdAt": "2020-07-02T18:12:51Z",
					"starsReceived": 0
				},
			]
	}`
	mockResponse := &services.UserScoreResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetUserScore(
		"login",
		"token",
		"provider",
	).Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/userscore", controller.GetUserScoreHandler)

	q := url.Values{}
	q.Add("login", "login")
	q.Add("access_token", "token")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/userscore?"+q.Encode(), nil)
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

	theContract := &services.UserScoreResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
