package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationRule interface {
	Init(rule string) error
	Apply(value reflect.Value) error
}

type intMinRule struct {
	min int
}

type strLenRule struct {
	len int
}

func (s *strLenRule) Init(rule string) error {
	length, err := strconv.Atoi(rule)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidValidationRule, err)
	}
	if length < 0 {
		return fmt.Errorf("%w: length should'n be negative", ErrInvalidValidationRule)
	}
	s.len = length
	return nil
}

func (s *strLenRule) Apply(value reflect.Value) error {
	v := value.String()
	length := len(v)
	switch {
	case length < s.len:
		return fmt.Errorf(`%w: "%s"`, ErrTooShortString, v)
	case length > s.len:
		return fmt.Errorf(`%w: "%s"`, ErrTooLongString, v)
	}
	return nil
}

type strRegexpRule struct {
	pattern *regexp.Regexp
}

func (s *strRegexpRule) Init(rule string) error {
	pattern, err := regexp.Compile(rule)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidValidationRule, err)
	}
	s.pattern = pattern
	return nil
}

func (s *strRegexpRule) Apply(value reflect.Value) error {
	if !s.pattern.MatchString(value.String()) {
		return ErrInvalidPattern
	}
	return nil
}

type strInRule struct {
	in []string
}

func (s *strInRule) Init(rule string) error {
	s.in = strings.Split(rule, ",")
	return nil
}

func (s *strInRule) Apply(value reflect.Value) error {
	if v := value.String(); !containsString(s.in, v) {
		return fmt.Errorf(`%w: actual: "%s", should be in %v`, ErrUnknownValue, v, s.in)
	}
	return nil
}

func (i *intMinRule) Init(str string) error {
	min, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidValidationRule, err)
	}
	i.min = min
	return nil
}

func (i *intMinRule) Apply(value reflect.Value) error {
	if v := int(value.Int()); v < i.min {
		return fmt.Errorf("%w: actual: %d, should be >= %d", ErrTooSmallNumber, v, i.min)
	}
	return nil
}

type intMaxRule struct {
	max int
}

func (i *intMaxRule) Init(str string) error {
	max, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidValidationRule, err)
	}
	i.max = max
	return nil
}

func (i *intMaxRule) Apply(value reflect.Value) error {
	if v := int(value.Int()); v > i.max {
		return fmt.Errorf("%w: actual: %d, should be <= %d", ErrTooBigNumber, v, i.max)
	}
	return nil
}

type intInRule struct {
	in []int
}

func (i *intInRule) Init(str string) error {
	parts := strings.Split(str, ",")
	in := make([]int, 0, len(parts))
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidValidationRule, err)
		}
		in = append(in, i)
	}
	i.in = in
	return nil
}

func (i *intInRule) Apply(value reflect.Value) error {
	if v := int(value.Int()); !containsInt(i.in, v) {
		return fmt.Errorf("%w: actual: %d, should be in %v", ErrUnknownValue, v, i.in)
	}
	return nil
}
