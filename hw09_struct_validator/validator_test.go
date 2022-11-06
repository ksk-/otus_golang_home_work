package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

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

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	Person struct {
		Name string `validate:"****"` // invalid rule syntax
	}

	Record struct {
		ID string `validate:"rule:value"` // unknown rule
	}

	Point struct {
		X int `validate:"min:0"`
		Y int `validate:"min:0|max:100|min:1"` // duplicate rule
	}

	Position struct {
		X int `validate:"min:0|max:100"`
		Y int `validate:"min:0|max:100"`
	}

	Rectangle struct {
		Pos    Position `validate:"nested"`
		Width  int      `validate:"min:0|max:100"`
		Height int      `validate:"min:0|max:100"`
	}

	Label struct {
		App       App
		Rect      Rectangle `validate:"nested"`
		Name      string    `validate:"in:foo,bar"`
		Timestamp time.Time
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			User{
				ID:     "user_1",
				Name:   "user 1",
				Age:    78,
				Email:  "invalid.email",
				Role:   "unknown",
				Phones: []string{"011-022-345", "+1-011-022-345"},
				meta:   nil,
			},
			ValidationErrors{
				{"ID", ErrTooShortString},
				{"Age", ErrTooBigNumber},
				{"Email", ErrInvalidPattern},
				{"Role", ErrUnknownValue},
				{"Phones", ErrTooLongString},
			},
		},

		{App{Version: "0.1"}, ValidationErrors{{"Version", ErrTooShortString}}},
		{App{Version: "0.1.2.3"}, ValidationErrors{{"Version", ErrTooLongString}}},
		{App{Version: "0.1.2"}, nil},

		{Token{Header: []byte("h"), Payload: []byte("p"), Signature: []byte("s")}, nil},

		{Response{Code: 201, Body: "{}"}, ValidationErrors{{"Code", ErrUnknownValue}}},
		{Response{Code: 200, Body: "{}"}, nil},

		{Person{Name: "Alice"}, ErrInvalidValidationRule},
		{Record{ID: "**"}, ErrUnknownValidationRule},
		{Point{X: 1, Y: 2}, ErrDuplicateValidationRule},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			requireEqualErrors(t, tt.expectedErr, Validate(tt.in))
		})
	}

	t.Run("no structs", func(t *testing.T) {
		for _, v := range []interface{}{1, -3.14, "text", true, func() {}} {
			require.ErrorIs(t, Validate(v), ErrNoStruct)
		}
	})
}

func TestValidate_WithNestedStruct(t *testing.T) {
	label := Label{
		App:       App{Version: "0.1"},
		Rect:      Rectangle{Pos: Position{X: -11, Y: 100}, Width: 10, Height: 111},
		Name:      "name",
		Timestamp: time.Now(),
	}

	expected := ValidationErrors{
		{
			"Rect", ValidationErrors{
				{"Pos", ValidationErrors{{"X", ErrTooSmallNumber}}},
				{"Height", ErrTooBigNumber},
			},
		},
		{"Name", ErrUnknownValue},
	}
	requireEqualErrors(t, expected, Validate(label))
}

func requireEqualErrors(t *testing.T, expected error, actual error) {
	t.Helper()

	if expected == nil {
		require.NoError(t, actual)
		return
	}

	var expectedErrs ValidationErrors
	if errors.As(expected, &expectedErrs) {
		var actualErrs ValidationErrors
		require.ErrorAs(t, actual, &actualErrs)

		require.Equal(t, len(expectedErrs), len(actualErrs))
		for i := 0; i < len(expectedErrs); i++ {
			require.Equal(t, expectedErrs[i].Field, actualErrs[i].Field)
			requireEqualErrors(t, expectedErrs[i].Err, actualErrs[i].Err)
		}
	} else {
		require.ErrorIs(t, actual, expected)
	}
}
