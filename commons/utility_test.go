package commons

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	asserts := assert.New(t)

	new_error := NewError("database", errors.New("invalid query"))
	err, ok := new_error.Errors["database"]

	asserts.Equal("invalid query", err)
	asserts.True(ok)
}

func TestRandString(t *testing.T) {
	asserts := assert.New(t)

	str := RandString(0)
	asserts.Equal(str, "", "Length should be 0")

	str = RandString(10)
	asserts.Equal(len(str), 10, "length should be 10")
	// fmt.Println("str", str)

	//check for containing letter and number
	contain_letter, contain_number := false, false
	for _, ch := range str {
		if !contain_letter {
			contain_letter = isLetter(string(ch))
		}
		if !contain_number {
			contain_number = isNumber(string(ch))
		}
	}
	asserts.True(contain_letter, "string should contain letter")
	asserts.True(contain_number, "string should contain number")

	time.Sleep(time.Millisecond * 100) //delay to ensure different seed value

	//generate another string
	str1 := RandString(10)
	asserts.Equal(len(str), 10, "length should be 10")
	asserts.NotEqual(str, str1, "string should not equal")
	// fmt.Println("str1", str1)
}
