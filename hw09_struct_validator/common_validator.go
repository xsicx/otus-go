package hw09structvalidator

import (
	"fmt"
)

type comparedValue interface {
	int | string
}

func validateIn[T comparedValue](v T, iNs []T) *ValidationError {
	for _, in := range iNs {
		if in == v {
			return nil
		}
	}

	return &ValidationError{Err: fmt.Errorf("value must contains in %v", iNs)}
}
