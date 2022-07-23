package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrInvalidValidationRule = errors.New("invalid validation rule")
	ErrInvalidStruct         = errors.New("invalid struct")
)

type ValidationError struct {
	Field string
	Err   error
}

type Rule struct {
	name      string
	condition string
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	builder := strings.Builder{}
	for _, err := range v {
		builder.WriteString(fmt.Sprintf("%s: %v\n", err.Field, err.Err))
	}

	return builder.String()
}

func Validate(v interface{}) error {
	var validationErrs ValidationErrors
	rValue := reflect.Indirect(reflect.ValueOf(v))

	if rValue.Kind() != reflect.Struct {
		return ErrInvalidStruct
	}

	for i := 0; i < rValue.NumField(); i++ {
		fieldType := rValue.Type().Field(i)
		validationRules, exists := fieldType.Tag.Lookup("validate")
		if !exists {
			continue
		}
		fieldValue := rValue.Field(i)

		rulesForField := strings.Split(validationRules, "|")

		for _, ruleForField := range rulesForField {
			var validateErr *ValidationError
			var err error
			ruleParts := strings.Split(ruleForField, ":")
			if len(ruleParts) != 2 {
				return ErrInvalidValidationRule
			}
			rule := Rule{
				name:      ruleParts[0],
				condition: ruleParts[1],
			}
			validateErr, err = validateKind(fieldValue, rule)

			if err != nil {
				return err
			}

			if validateErr != nil {
				validateErr.Field = fieldType.Name
				validationErrs = append(validationErrs, *validateErr)
			}
		}
	}

	if len(validationErrs) != 0 {
		return validationErrs
	}

	return nil
}

func validateKind(fieldValue reflect.Value, rule Rule) (*ValidationError, error) {
	switch fieldValue.Kind() { // nolint: exhaustive
	case reflect.String:
		return validateString(fieldValue.String(), rule)
	case reflect.Int:
		return validateInt(int(fieldValue.Int()), rule)
	case reflect.Slice:
		switch fieldValue.Interface().(type) {
		case []string:
			for _, val := range fieldValue.Interface().([]string) {
				validateErr, err := validateString(val, rule)
				if err != nil || validateErr != nil {
					return validateErr, err
				}
			}
		case []int:
			for _, val := range fieldValue.Interface().([]int) {
				validateErr, err := validateInt(val, rule)
				if err != nil || validateErr != nil {
					return validateErr, err
				}
			}
		}
	}

	return nil, nil
}
