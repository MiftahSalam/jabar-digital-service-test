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

var MockRegister = []RouterMockTest{
	{
		TestName: "no error: Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"%v","role":"%v"}`,
			model.UsersMock[0].Username,
			model.UsersMock[0].UserRole.Name,
		),
		ResponseCode:    http.StatusCreated,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusOK",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			var jsonResp service.RegisteredUserResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				fmt.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// fmt.Println("jsonResp", jsonResp)

			a.Equal(model.UsersMock[0].Username, jsonResp.User.Username)
			a.Equal(model.UsersMock[0].UserRole.Name, jsonResp.User.Role)
			a.NotEmpty(jsonResp.User.Password)
		},
	},
	{
		TestName: "error bad request (no data body) : Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "/auth/",
		Method:          "POST",
		Body:            "",
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "error")
		},
	},
	{
		TestName: "error bad request (empty data body) : Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "/auth/",
		Method:          "POST",
		Body:            "{}",
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "error")
			a.Equal(2, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		TestName: "error bad request (no username data body) : Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"","role":"%v"}`,
			model.UsersMock[0].UserRole.Name,
		),
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "error")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		TestName: "error bad request (no role data body) : Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"%v","role":""}`,
			model.UsersMock[0].Username,
		),
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "error")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		TestName: "	error (role does not exist): Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"%v","role":"MyRole"}`,
			model.UsersMock[0].Username,
		),
		ResponseCode:    http.StatusBadRequest,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusBadRequest",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "invalid role")
		},
	},
	{
		TestName: "	error unprocessed data (user already exist): Register Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:    "/auth/",
		Method: "POST",
		Body: fmt.Sprintf(`{"username":"%v","role":"%v"}`,
			model.UsersMock[0].Username,
			model.UsersMock[0].UserRole.Name,
		),
		ResponseCode:    http.StatusUnprocessableEntity,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusUnprocessableEntity",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			a.Contains(string(response_body), "user already exist")
		},
	},
}
