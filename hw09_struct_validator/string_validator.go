package hw09structvalidator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validateString(v string, rule Rule) (*ValidationError, error) {
	switch rule.name {
	case "len":
		strLen, err := strconv.ParseInt(rule.condition, 10, 0)
		if err != nil {
			return nil, ErrInvalidValidationRule
		}
		return validateLen(v, int(strLen)), nil
	case "in":
		iNs := strings.Split(rule.condition, ",")
		return validateIn(v, iNs), nil
	case "regexp":
		regex, err := regexp.Compile(rule.condition)
		if err != nil {
			return nil, ErrInvalidValidationRule
		}
		return validateRegex(v, regex), nil
	}

	return nil, nil
}

func validateLen(v string, strLen int) *ValidationError {
	if len(v) != strLen {
		return &ValidationError{Err: fmt.Errorf("size of value must be %d", strLen)}
	}

	return nil
}

func validateRegex(v string, regexp *regexp.Regexp) *ValidationError {
	if !regexp.MatchString(v) {
		return &ValidationError{Err: fmt.Errorf("value must match by '%s'", regexp.String())}
	}

	return nil
}
