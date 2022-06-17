package router

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/MiftahSalam/jabar-digital-service-test/users/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Test main users router start")

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Cannot load env file. Err: ", err)
		panic("Cannot load env file")
	}

	err = database.InitDatabase()
	if err != nil {
		panic("Database error connection")
	}

	model.Migrate()
	validator.Init()

	gin.SetMode(gin.TestMode)

	router = gin.New()
	Users(router.Group("/auth"))

	exitVal := m.Run()

	model.CleaningUpTest()
	database.CloseDatabase()

	os.Exit(exitVal)

	fmt.Println("Test main users router end")
}

func createTest(asserts *assert.Assertions, testData *RouterMockTest) *httptest.ResponseRecorder {
	body := testData.Body
	req, err := http.NewRequest(testData.Method, testData.Url, bytes.NewBufferString(body))

	fmt.Println("test url", testData.Url)
	// common.LogI.Println("test body", testData.UserMockTest.Body)

	asserts.NoError(err)

	testData.Init(req)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	asserts.Equal(testData.ResponseCode, w.Code, "Response Status - "+testData.Msg)

	return w
}

func TestRegister(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockRegister {
		t.Run(test.TestName, func(t *testing.T) {
			w := createTest(asserts, &test)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(w, asserts)
		})
	}
}

func TestLogin(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockLogin {
		t.Run(test.TestName, func(t *testing.T) {
			if strings.Contains(test.TestName, "no error") {
				test.Body = fmt.Sprintf(`{"username":"%v","password":"%v"}`,
					model.UsersMock[0].Username,
					registeredPassword,
				)
			}
			w := createTest(asserts, &test)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(w, asserts)
		})
	}
}

func TestValidate(t *testing.T) {
	asserts := assert.New(t)

	for _, test := range MockValidate {
		t.Run(test.TestName, func(t *testing.T) {
			if strings.Contains(test.TestName, "no error") {
				test.Url = fmt.Sprintf("/auth/%v", LoggedInToken)
			} else if strings.Contains(test.TestName, "error not authorized (signature invalid)") {
				test.Url = fmt.Sprintf("/auth/%v", LoggedInToken[:len(LoggedInToken)-1])
			}
			w := createTest(asserts, &test)

			asserts.Equal(test.ResponseCode, w.Code)

			test.ResponseTest(w, asserts)
		})
	}
}
