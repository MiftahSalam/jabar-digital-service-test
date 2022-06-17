package router

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type RouterMockTest struct {
	TestName        string
	Init            func(*http.Request)
	Url             string
	Method          string
	Body            string
	ResponseCode    int
	ResponsePattern string
	Msg             string
	ResponseTest    func(w *httptest.ResponseRecorder, a *assert.Assertions)
}

var router *gin.Engine
