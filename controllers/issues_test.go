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
	"github.com/golang/mock/gomock"
	"github.com/labbsr0x/githunter-api/services"
	"github.com/labbsr0x/githunter-api/services/mock"
)

func TestIssuesController_GetIssuesHandler_Error_GetIssues_Unknown_Provider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	// Mocking the values Expected
	mockContractService.EXPECT().GetIssues(10, "owner", "repo", "provider", "accessToken").Return(nil, fmt.Errorf("GetIssues unknown provider: provider"))

	app := fiber.New()
	app.Get("/issues", controller.GetIssuesHandler)

	q := url.Values{}
	q.Add("owner", "owner")
	q.Add("name", "repo")
	q.Add("access_token", "accessToken")
	q.Add("provider", "provider")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/issues?"+q.Encode(), nil)
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

func TestIssuesController_GetIssuesHandler_Error_GetIssues_Invalid_Token(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	// Mocking the values Expected
	mockContractService.EXPECT().GetIssues(10, "owner", "repo", "provider", "invalidToken").Return(nil, fmt.Errorf("'Get' using unknown token: invalidToken"))

	app := fiber.New()
	app.Get("/issues", controller.GetIssuesHandler)

	q := url.Values{}
	q.Add("access_token", "invalidToken")
	q.Add("provider", "provider")
	q.Add("owner", "owner")
	q.Add("name", "repo")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/issues?"+q.Encode(), nil)
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

func TestIssuesController_GetIssuesHandler_Error_GetIssues_Invalid_NameAndOwner(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockContractService := mock.NewMockContract(mockController)

	controller := &Controller{
		Contract: mockContractService,
	}

	// Mocking the values Expected
	mockContractService.EXPECT().GetPulls(10, "invalidOwner", "invalidRepo", "provider", "validToken").Return(nil, nil)

	app := fiber.New()
	app.Get("/issues", controller.GetPullsHandler)

	q := url.Values{}
	q.Add("access_token", "validToken")
	q.Add("provider", "provider")
	q.Add("owner", "invalidOwner")
	q.Add("name", "invalidRepo")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/issues?"+q.Encode(), nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// http.Response
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusNoContent {
		t.Errorf("expect response status code %d got %d", fiber.StatusNoContent, resp.StatusCode)
	}

}

func TestIssuesController_GetIssuesHandler_Success(t *testing.T) {
	mockContractService, controller := GetMockContractServiceAndController(t)

	responseJSONStr := `{
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
			{
				"number": 12794,
				"state": "OPEN",
				"createdAt": "2020-08-06T07:45:36Z",
				"updatedAt": "2020-08-07T10:39:59Z",
				"closedAt": "",
				"author": "fakeAuthor",
				"labels": [
					"fakeLabel"
				],
				"totalParticipants": 2,
				"timelineItems": {
					"totalCount": 1,
					"updatedAt": "2020-08-06T13:04:34Z",
					"nodes": [
						{
							"__typename": "IssueComment",
							"createdAt": "2020-08-06T13:04:34Z",
							"author": "fakeAuthor"
						}
					]
				}
			}`

	mockResponse := &services.IssuesResponseContract{}
	json.Unmarshal([]byte(responseJSONStr), mockResponse)

	// Mocking the values Expected
	mockContractService.EXPECT().GetIssues(10, "validOwner", "validName", "github", "validToken").Return(mockResponse, nil)

	app := fiber.New()
	app.Get("/issues", controller.GetIssuesHandler)

	q := url.Values{}
	q.Add("access_token", "validToken")
	q.Add("provider", "github")
	q.Add("owner", "validOwner")
	q.Add("name", "validName")

	// http.Request
	req := httptest.NewRequest(fiber.MethodGet, "/issues?"+q.Encode(), nil)
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

	theContract := &services.IssuesResponseContract{}
	json.Unmarshal(respBody, theContract)

	if !reflect.DeepEqual(theContract, mockResponse) {
		t.Errorf("expect %#v response message. got %#v", mockResponse, theContract)
	}
}
