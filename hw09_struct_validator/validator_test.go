package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	User2 struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	BrokenRule struct {
		Value string `validate:"len:five"`
	}
	BrokenRule2 struct {
		Value string `validate:""`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
		weight    float64 `validate:"min:5"`
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		// app errors case
		{[]int{1, 2}, ErrInvalidStruct},
		{BrokenRule{Value: "test"}, ErrInvalidValidationRule},
		{BrokenRule2{Value: "test"}, ErrInvalidValidationRule},
		// without validation rules case
		{Token{
			Header:    []byte{},
			Payload:   []byte{},
			Signature: []byte{},
			weight:    1,
		}, nil},

		{App{Version: "v1234"}, nil},
		{User{
			ID:     "4f0e8c90-2854-49bb-ab29-faf72f41be06",
			Name:   "Ivan",
			Age:    20,
			Email:  "ivan@mail.ru",
			Role:   "stuff",
			Phones: []string{"n1234567890", "n0987654321"},
			meta:   []byte{},
		}, nil},

		{
			Response{
				Code: 301,
				Body: "{}",
			}, ValidationErrors{
				{
					Field: "Code",
					Err:   errors.New(`value must contains in [200 404 500]`),
				},
			},
		},
		{
			User{
				ID:     "13",
				Name:   "Ivan",
				Age:    10,
				Email:  "sdfdf",
				Role:   "user",
				Phones: []string{"n123456890", "n434"},
			}, ValidationErrors{
				{
					Field: "ID",
					Err:   errors.New(`size of value must be 36`),
				},
				{
					Field: "Age",
					Err:   errors.New(`value must be no less than 18`),
				},
				{
					Field: "Email",
					Err:   fmt.Errorf("value must match by '%s'", "^\\w+@\\w+\\.\\w+$"),
				},
				{
					Field: "Role",
					Err:   errors.New("value must contains in [admin stuff]"),
				},
				{
					Field: "Phones",
					Err:   errors.New("size of value must be 11"),
				},
			},
		},
		// app error with validations error case
		{
			User2{
				ID:     "13",
				Name:   "Ivan",
				Age:    10,
				Email:  "sdfdf",
				Role:   "user",
				Phones: []string{"n123456890", "n434"},
			}, ErrInvalidValidationRule,
		},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.expectedErr, Validate(tt.in))
		})
	}
}
