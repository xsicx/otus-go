package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var strBuilder strings.Builder
	lastSymbol, isEscaped := "", false

	if !isCorrectString(str) {
		return "", ErrInvalidString
	}

	for _, val := range str {
		if val == '\\' && !isEscaped {
			isEscaped = true
			continue
		}

		if unicode.IsDigit(val) && !isEscaped {
			if val == '0' {
				removeLastSymbol(&strBuilder)
				continue
			}
			size, _ := strconv.Atoi(string(val))
			strBuilder.WriteString(strings.Repeat(lastSymbol, size-1))
		} else {
			strBuilder.WriteRune(val)
			isEscaped = false
		}

		lastSymbol = string(val)
	}

	return strBuilder.String(), nil
}

func isCorrectString(str string) bool {
	prevSymbolIsDigit, isEscaped := false, false

	for i, val := range str {
		if val == '\\' && !isEscaped {
			isEscaped = true
			continue
		}

		if unicode.IsDigit(val) && !isEscaped {
			if i == 0 || prevSymbolIsDigit {
				return false
			}

			prevSymbolIsDigit = true
			continue
		}

		isEscaped = false
		prevSymbolIsDigit = false
	}

	return true
}

func removeLastSymbol(builder *strings.Builder) {
	str := builder.String()
	builder.Reset()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	builder.WriteString(str)
}
