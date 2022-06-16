package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	db "github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/MiftahSalam/jabar-digital-service-test/users/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Test User Model package start")

	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Cannot load environment file")
	}

	err = db.InitDatabase()
	if err != nil {
		panic("Database error connection")
	}
	model.Migrate()

	validator.Init()

	gin.SetMode(gin.TestMode)

	exit_code := m.Run()

	model.CleaningUpTest()
	db.CloseDatabase()
	os.Exit(exit_code)
	fmt.Println("Test User Model package end")
}

func InitTest() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	return c, w
}

func MockJSONPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbyte, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Cannot marshal json content")
	}
	fmt.Println("MockJSONPost req content", string(jsonbyte))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbyte))
}

func TestRegister(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockRegisterTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)
			MockJSONPost(c, test.Data)

			Register(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}

	// user, _ := c.Get("user")
	// common.LogI.Println("key user before", user)
	// for k := range c.Keys {
	// 	if k == "user" {
	// 		delete(c.Keys, k)
	// 		common.LogI.Println("ckeys", k)
	// 	}
	// }
	// user, _ = c.Get("user")
	// common.LogI.Println("key user after", user)
}
