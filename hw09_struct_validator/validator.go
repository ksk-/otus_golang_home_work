package hw09structvalidator

import (
	"fmt"
	"reflect"
)

const validateTag = "validate"

type ValidationError struct {
	Field string
	Err   error
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(`field: "%s" validation error: %v`, v.Field, v.Err)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errors := make([]string, len(v))
	for i, vErr := range v {
		errors[i] = fmt.Sprintf("'%s': %v", vErr.Field, vErr.Err)
	}
	return fmt.Sprintf("%v", errors)
}

func Validate(v interface{}) error {
	rValue := reflect.ValueOf(v)
	rType := reflect.TypeOf(v)
	if rType.Kind() != reflect.Struct {
		return ErrNoStruct
	}

	var validationErrors ValidationErrors
	for i := 0; i < rValue.NumField(); i++ {
		field := rType.Field(i)
		if tag, ok := field.Tag.Lookup(validateTag); ok {
			fv, err := NewFieldValidator(field.Type.Kind(), tag)
			if err != nil {
				return fmt.Errorf("create field validator: %w", err)
			}
			if fv == nil {
				continue
			}

			if err = fv.Validate(rValue.Field(i)); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
