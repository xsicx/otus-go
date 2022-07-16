package hw09structvalidator

import (
	"fmt"
	"strconv"
	"strings"
)

func validateInt(v int, rule Rule) (*ValidationError, error) {
	switch rule.name {
	case "min":
		minVal, err := strconv.ParseInt(rule.condition, 10, 0)
		if err != nil {
			return nil, ErrInvalidValidationRule
		}
		return validateMin(v, int(minVal)), nil
	case "max":
		maxVal, err := strconv.ParseInt(rule.condition, 10, 0)
		if err != nil {
			return nil, ErrInvalidValidationRule
		}
		return validateMax(v, int(maxVal)), nil
	case "in":
		var iNs []int
		iNsStr := strings.Split(rule.condition, ",")
		for _, inStr := range iNsStr {
			in, err := strconv.ParseInt(inStr, 10, 0)
			if err != nil {
				return nil, ErrInvalidValidationRule
			}
			iNs = append(iNs, int(in))
		}
		return validateIn(v, iNs), nil
	}

	return nil, nil
}

func validateMin(v int, minVal int) *ValidationError {
	if v < minVal {
		return &ValidationError{Err: fmt.Errorf("value must be no less than %d", minVal)}
	}

	return nil
}

func validateMax(v int, maxVal int) *ValidationError {
	if v > maxVal {
		return &ValidationError{Err: fmt.Errorf("value must be no greater than %d", maxVal)}
	}

	return nil
}
