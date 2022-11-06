package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type FieldValidator interface {
	Validate(value reflect.Value) error
}

func NewFieldValidator(kind reflect.Kind, tag string) (FieldValidator, error) {
	if kind == reflect.String {
		return newScalarFieldValidator(tag, map[string]ValidationRule{
			"len":    &strLenRule{},
			"regexp": &strRegexpRule{},
			"in":     &strInRule{},
		})
	}
	if kind == reflect.Int {
		return newScalarFieldValidator(tag, map[string]ValidationRule{
			"min": &intMinRule{},
			"max": &intMaxRule{},
			"in":  &intInRule{},
		})
	}
	if kind == reflect.Slice {
		return &sliceValidator{tag}, nil
	}
	if kind == reflect.Struct {
		return newStructValidator(tag)
	}
	return nil, nil
}

type scalarFieldValidator struct {
	rules map[string]ValidationRule
}

func (s *scalarFieldValidator) Validate(value reflect.Value) error {
	for _, rule := range s.rules {
		if err := rule.Apply(value); err != nil {
			return fmt.Errorf("invalid field: %w", err)
		}
	}
	return nil
}

func (s *scalarFieldValidator) addRule(name string, rule ValidationRule) error {
	if _, ok := s.rules[name]; ok {
		return ErrDuplicateValidationRule
	}
	s.rules[name] = rule
	return nil
}

func newScalarFieldValidator(tag string, rules map[string]ValidationRule) (*scalarFieldValidator, error) {
	v := &scalarFieldValidator{
		rules: make(map[string]ValidationRule),
	}

	re := regexp.MustCompile(`^(\w+):(.*)$`)
	for _, r := range strings.Split(tag, "|") {
		if !re.MatchString(r) {
			return nil, ErrInvalidValidationRule
		}

		groups := re.FindStringSubmatch(r)
		name := groups[1]
		rule, ok := rules[name]
		if !ok {
			return nil, ErrUnknownValidationRule
		}
		if err := rule.Init(groups[2]); err != nil {
			return nil, fmt.Errorf("rule init: %w", err)
		}
		if err := v.addRule(name, rule); err != nil {
			return nil, fmt.Errorf("validator add rule: %w", err)
		}
	}

	return v, nil
}

type sliceValidator struct {
	tag string
}

func (s *sliceValidator) Validate(value reflect.Value) error {
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		v, err := NewFieldValidator(item.Kind(), s.tag)
		if err != nil {
			return fmt.Errorf("slice item #%d validator create: %w", i, err)
		}
		if err = v.Validate(item); err != nil {
			return fmt.Errorf("slice item #%d validate: %w", i, err)
		}
	}
	return nil
}

type structValidator struct{}

func (s *structValidator) Validate(value reflect.Value) error {
	return Validate(value.Interface())
}

func newStructValidator(tag string) (*structValidator, error) {
	if tag != "nested" {
		return nil, ErrInvalidValidationRule
	}
	return &structValidator{}, nil
}
