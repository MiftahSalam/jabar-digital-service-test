package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/MiftahSalam/jabar-digital-service-test/users/service"
	"github.com/stretchr/testify/assert"
)

var MockValidate = []RouterMockTest{
	{
		TestName: "no error: Validate Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "",
		Method:          "GET",
		Body:            "",
		ResponseCode:    http.StatusOK,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusOK",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			var jsonResp service.ValidateTokenResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				fmt.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// fmt.Println("jsonResp", jsonResp)

			a.Equal(model.UsersMock[0].Username, jsonResp.User.Username)
			a.True(jsonResp.User.Is_valid)
		},
	},
	{
		TestName: "error not authorized (signature invalid): Auth Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "",
		Method:          "GET",
		Body:            "",
		ResponseCode:    http.StatusUnauthorized,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusUnauthorized",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"auth":"signature is invalid"}}`, string(response_body))
		},
	},
	{
		TestName: "error not authorized (invalid token): Auth Test",
		Init: func(req *http.Request) {
			req.Header.Set("Content-Type", "application/json")
		},
		Url:             "/auth/d",
		Method:          "GET",
		Body:            "",
		ResponseCode:    http.StatusUnauthorized,
		ResponsePattern: "",
		Msg:             "valid data and should return StatusUnauthorized",
		ResponseTest: func(w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			a.Equal(`{"errors":{"auth":"token contains an invalid number of segments"}}`, string(response_body))
		},
	},
}
