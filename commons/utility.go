package commons

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
	new_error := CommonError{}
	new_error.Errors = make(map[string]interface{})
	new_error.Errors[key] = err.Error()

	return new_error
}

func NewValidationError(err error) CommonError {
	new_error := CommonError{}
	new_error.Errors = make(map[string]interface{})
	var current_error *json.SyntaxError

	if errors.As(err, &current_error) {
		new_error.Errors["json_error"] = err.Error()
		return new_error
	}

	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param() != "" {
			new_error.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			new_error.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}

	return new_error
}

var number = []rune("0123456789")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphanum = append(letters, number...)

func isNumber(ch string) bool {
	return strings.Contains(string(number), ch)
}

func isLetter(ch string) bool {
	return strings.Contains(string(letters), ch)
}

func RandString(n int) string {
	if n <= 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	alphanum_len := len(alphanum)
	contain_letter, contain_number := false, false

	for i := range b {
		b[i] = alphanum[rand.Intn(alphanum_len)]
		if !contain_letter {
			contain_letter = isLetter(string(b[i]))
		}
		if !contain_number {
			contain_number = isNumber(string(b[i]))
		}
	}

	if !contain_letter {
		b[rand.Intn(n)] = letters[rand.Intn(len(letters))]
	} else if !contain_number {
		b[rand.Intn(n)] = number[rand.Intn(len(number))]
	}
	return string(b)

}