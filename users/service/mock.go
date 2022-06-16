package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/MiftahSalam/jabar-digital-service-test/users/serializer"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockTests struct {
	TestName     string
	Init         func(c *gin.Context)
	Data         interface{}
	ResponseCode int
	ResponseTest func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions)
}
type RegisteredUserResponse struct {
	User struct {
		serializer.UserRegisteredResponse
	} `json:"user"`
}

var MockRegisterTest = []MockTests{
	{
		"no error: Register Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"username": model.UsersMock[0].Username,
			"role":     model.UsersMock[0].UserRole.Name,
		},
		http.StatusCreated,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))

			var jsonResp RegisteredUserResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				fmt.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			fmt.Println("jsonResp", jsonResp)

			a.Equal(model.UsersMock[0].Username, jsonResp.User.Username)
			a.Equal(model.UsersMock[0].UserRole.Name, jsonResp.User.Role)
			a.Equal(6, len(jsonResp.User.Password))
		},
	},
	{
		"error bad request (no username data body): Register Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"role": model.UsersMock[0].UserRole.Name,
		},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			fmt.Println("response_body", string(response_body))
			a.Contains(string(response_body), "errors")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		"error bad request (no role data body): Register Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"username": model.UsersMock[0].Username,
		},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			fmt.Println("response_body", string(response_body))
			a.Contains(string(response_body), "errors")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		"error bad request (no data body): Register Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			fmt.Println("response_body", string(response_body))
			a.Contains(string(response_body), "errors")
			a.Equal(2, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		"error (user already exist): Register Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"username": model.UsersMock[0].Username,
			"role":     model.UsersMock[0].UserRole.Name,
		},
		http.StatusUnprocessableEntity,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			fmt.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"database":"user already exist"}}`, string(response_body))
		},
	},
}
