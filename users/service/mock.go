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
type LoggedInUserResponse struct {
	User struct {
		serializer.UserLoggedInResponse
	} `json:"user"`
}

var LoggedInPassword string
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

			// fmt.Println("response_body", string(response_body))

			var jsonResp RegisteredUserResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				fmt.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// fmt.Println("jsonResp", jsonResp)

			a.Equal(model.UsersMock[0].Username, jsonResp.User.Username)
			a.Equal(model.UsersMock[0].UserRole.Name, jsonResp.User.Role)
			a.Equal(6, len(jsonResp.User.Password))
			LoggedInPassword = jsonResp.User.Password
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
			// fmt.Println("response_body", string(response_body))
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
			// fmt.Println("response_body", string(response_body))
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
			// fmt.Println("response_body", string(response_body))
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

			// fmt.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"database":"user already exist"}}`, string(response_body))
		},
	},
}

var MockLoginTest = []MockTests{
	{
		"no error: Login Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{},
		http.StatusOK,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)

			// fmt.Println("response_body", string(response_body))

			var jsonResp LoggedInUserResponse
			err := json.Unmarshal(response_body, &jsonResp)
			if err != nil {
				fmt.Println("Cannot umarshal json content with error: ", err)
			}
			a.NoError(err)

			// fmt.Println("jsonResp", jsonResp)

			a.Equal(model.UsersMock[0].Username, jsonResp.User.Username)
			a.NotEmpty(jsonResp.User.Token, "token should not empty string")
		},
	},
	{
		"error bad request (no username data body): Login Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"password": "124567",
		},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// fmt.Println("response_body", string(response_body))
			a.Contains(string(response_body), "errors")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		"error bad request (no password data body): Login Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"username": model.UsersMock[0].Username,
		},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// fmt.Println("response_body", string(response_body))
			a.Contains(string(response_body), "errors")
			a.Equal(1, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		"error bad request (no data body): Login Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{},
		http.StatusBadRequest,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// fmt.Println("response_body", string(response_body))
			a.Contains(string(response_body), "errors")
			a.Equal(2, strings.Count(string(response_body), "key: required"))
		},
	},
	{
		"error not found (user not found): Login Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"username": "nouser",
			"password": "123456",
		},
		http.StatusNotFound,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			// fmt.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"database":"user not found"}}`, string(response_body))
		},
	},
	{
		"error forbidden (invalid password): Login Test",
		func(c *gin.Context) {
		},
		map[string]interface{}{
			"username": model.UsersMock[0].Username,
			"password": "123456",
		},
		http.StatusForbidden,
		func(c *gin.Context, w *httptest.ResponseRecorder, a *assert.Assertions) {
			response_body, _ := ioutil.ReadAll(w.Body)
			fmt.Println("response_body", string(response_body))
			a.Equal(`{"errors":{"login":"invalid password"}}`, string(response_body))
		},
	},
}
