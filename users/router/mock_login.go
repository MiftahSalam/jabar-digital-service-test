package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/MiftahSalam/jabar-digital-service-test/users/service"
	"github.com/stretchr/testify/assert"
)

var LoggedInToken string
var MockLogin = []RouterMockTest{
	{
		TestName: "no error: Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "/auth/login",
		Method:          "POST",
		Body:            "",
		ResponseCode:    http.StatusOK,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusOK",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			var jsonResp service.LoggedInUserResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				fmt.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// fmt.Println("jsonResp", jsonResp)

			a.Equal(model.UsersMock[0].Username, jsonResp.User.Username)
			a.NotEmpty(jsonResp.User.Token)

			LoggedInToken = jsonResp.User.Token
		},
	},
	{
		TestName: "error bad request (no username data body): Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/login",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"","password":"%v"}`,
			"registeredPassword",
		),
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "errors")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		TestName: "error bad request (no username data body): Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/login",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"%v","password":"%v"}`,
			model.UsersMock[0].Username,
			registeredPassword,
		),
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "errors")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		TestName: "error bad request (no data body): Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "/auth/login",
		Method:          "POST",
		Body:            "",
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "errors")
		},
	},
	{
		TestName: "error bad request (empty data body): Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "/auth/login",
		Method:          "POST",
		Body:            "{}",
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "errors")
			a.Equal(2, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		TestName: "error not found (user not found): Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/login",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"user","password":"%v"}`,
			"registeredPassword",
		),
		ResponseCode:    http.StatusNotFound,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusNotFound",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"database":"user not found"}}`, string(response_body))
		},
	},
	{
		TestName: "error not authorized (invalid password): Login Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/login",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"%v","password":"%v"}`,
			model.UsersMock[0].Username,
			"registeredPassword",
		),
		ResponseCode:    http.StatusForbidden,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusForbidden",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"login":"invalid password"}}`, string(response_body))
		},
	},
}
