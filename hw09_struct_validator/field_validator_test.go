package hw09structvalidator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateString_Len(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.String, "len:12")
		require.NoError(t, err)

		require.ErrorIs(t, v.Validate(reflect.ValueOf("short")), ErrTooShortString)
		require.ErrorIs(t, v.Validate(reflect.ValueOf("too long string")), ErrTooLongString)
		require.Nil(t, v.Validate(reflect.ValueOf("valid string")))
	})

	t.Run("invalid value", func(t *testing.T) {
		for _, tag := range []string{"len:invalid", "len:-12"} {
			v, err := NewFieldValidator(reflect.String, tag)
			require.Nil(t, v)
			require.ErrorIs(t, err, ErrInvalidValidationRule)
		}
	})
}

func TestValidateString_Regexp(t *testing.T) {
	v, err := NewFieldValidator(reflect.String, `regexp:^\d+:\w+$`)
	require.NoError(t, err)

	require.Nil(t, v.Validate(reflect.ValueOf("34:text")))
	require.ErrorIs(t, v.Validate(reflect.ValueOf("invalid_pattern")), ErrInvalidPattern)
}

func TestValidateString_In(t *testing.T) {
	v, err := NewFieldValidator(reflect.String, "in:foo,bar")
	require.NoError(t, err)

	require.Nil(t, v.Validate(reflect.ValueOf("foo")))
	require.Nil(t, v.Validate(reflect.ValueOf("bar")))
	require.ErrorIs(t, v.Validate(reflect.ValueOf("baz")), ErrUnknownValue)
}

func TestValidateString_SeveralRules(t *testing.T) {
	v, err := NewFieldValidator(reflect.String, `regexp:\d+|len:3`)
	require.NoError(t, err)

	require.Nil(t, v.Validate(reflect.ValueOf("123")))
	require.Nil(t, v.Validate(reflect.ValueOf("000")))
	require.ErrorIs(t, v.Validate(reflect.ValueOf("xxx")), ErrInvalidPattern)
	require.ErrorIs(t, v.Validate(reflect.ValueOf("12")), ErrTooShortString)

	t.Run("not unique rules", func(t *testing.T) {
		v, err = NewFieldValidator(reflect.String, `len:4|regexp:\d+|len:3`)
		require.Nil(t, v)
		require.ErrorIs(t, err, ErrDuplicateValidationRule)
	})
}

func TestValidateInt_Min(t *testing.T) {
	t.Run("positive value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "min:12")
		require.NoError(t, err)

		require.ErrorIs(t, v.Validate(reflect.ValueOf(11)), ErrTooSmallNumber)
		require.Nil(t, v.Validate(reflect.ValueOf(12)))
		require.Nil(t, v.Validate(reflect.ValueOf(13)))
	})

	t.Run("negative value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "min:-12")
		require.NoError(t, err)

		require.ErrorIs(t, v.Validate(reflect.ValueOf(-13)), ErrTooSmallNumber)
		require.Nil(t, v.Validate(reflect.ValueOf(-12)))
		require.Nil(t, v.Validate(reflect.ValueOf(-11)))
	})

	t.Run("invalid value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "min:invalid")
		require.Nil(t, v)
		require.ErrorIs(t, err, ErrInvalidValidationRule)
	})
}

func TestValidateInt_Max(t *testing.T) {
	t.Run("positive value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "max:12")
		require.NoError(t, err)

		require.ErrorIs(t, v.Validate(reflect.ValueOf(13)), ErrTooBigNumber)
		require.Nil(t, v.Validate(reflect.ValueOf(12)))
		require.Nil(t, v.Validate(reflect.ValueOf(11)))
	})

	t.Run("negative value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "max:-12")
		require.NoError(t, err)

		require.ErrorIs(t, v.Validate(reflect.ValueOf(-11)), ErrTooBigNumber)
		require.Nil(t, v.Validate(reflect.ValueOf(-12)))
		require.Nil(t, v.Validate(reflect.ValueOf(-13)))
	})

	t.Run("invalid value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "max:invalid")
		require.Nil(t, v)
		require.ErrorIs(t, err, ErrInvalidValidationRule)
	})
}

func TestValidateInt_In(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "in:-23,15")
		require.NoError(t, err)

		require.Nil(t, v.Validate(reflect.ValueOf(-23)))
		require.Nil(t, v.Validate(reflect.ValueOf(15)))
		require.ErrorIs(t, v.Validate(reflect.ValueOf(42)), ErrUnknownValue)
	})

	t.Run("invalid value", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Int, "in:-23,invalid")
		require.Nil(t, v)
		require.ErrorIs(t, err, ErrInvalidValidationRule)
	})
}

func TestValidateInt_SeveralRules(t *testing.T) {
	v, err := NewFieldValidator(reflect.Int, `min:5|max:10`)
	require.NoError(t, err)

	require.Nil(t, v.Validate(reflect.ValueOf(5)))
	require.Nil(t, v.Validate(reflect.ValueOf(7)))
	require.Nil(t, v.Validate(reflect.ValueOf(10)))
	require.ErrorIs(t, v.Validate(reflect.ValueOf(4)), ErrTooSmallNumber)
	require.ErrorIs(t, v.Validate(reflect.ValueOf(11)), ErrTooBigNumber)

	t.Run("not unique rules", func(t *testing.T) {
		v, err = NewFieldValidator(reflect.Int, `min:4|min:3|max:5`)
		require.Nil(t, v)
		require.ErrorIs(t, err, ErrDuplicateValidationRule)
	})
}

func TestValidateSlice(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Slice, "len:3")
		require.NoError(t, err)

		require.Nil(t, v.Validate(reflect.ValueOf([]string{"123", "456", "789"})))
		require.ErrorIs(t, v.Validate(reflect.ValueOf([]string{"123", "****", "789"})), ErrTooLongString)
		require.ErrorIs(t, v.Validate(reflect.ValueOf([]string{"123", "456", "*"})), ErrTooShortString)
	})

	t.Run("int", func(t *testing.T) {
		v, err := NewFieldValidator(reflect.Slice, "min:3|max:5")
		require.NoError(t, err)
		require.Nil(t, v.Validate(reflect.ValueOf([]int{3, 4, 5})))
		require.ErrorIs(t, v.Validate(reflect.ValueOf([]int{30, 4, 5})), ErrTooBigNumber)
		require.ErrorIs(t, v.Validate(reflect.ValueOf([]int{3, 4, -5})), ErrTooSmallNumber)
	})
}
