package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	// fmt.Println("MockJSONPost req content", string(jsonbyte))

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
}

func TestLogin(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockLoginTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			if strings.Contains(test.TestName, "no error") {
				data := map[string]interface{}{
					"username": model.UsersMock[0].Username,
					"password": LoggedInPassword,
				}
				MockJSONPost(c, data)
			} else {
				MockJSONPost(c, test.Data)
			}

			Login(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}

func TestAuth(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockAuthTest {
		t.Run(test.TestName, func(t *testing.T) {
			c, w := InitTest()
			test.Init(c)

			MockJSONPost(c, test.Data)

			Auth(c)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(c, w, asserts)
		})
	}
}
