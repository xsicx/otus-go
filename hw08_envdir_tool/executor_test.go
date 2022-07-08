package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type dataSet struct {
	cmd          []string
	expectedCode int
}

func TestRunCmd(t *testing.T) {
	testCases := []dataSet{
		{
			cmd:          []string{"echo"},
			expectedCode: 0,
		},
		{
			cmd:          []string{"/bin/bash", "echo"},
			expectedCode: 126,
		},
		{
			cmd:          []string{"/bin/bash", "command_not_found"},
			expectedCode: 127,
		},
	}
	t.Run("main test", func(t *testing.T) {
		for _, testCase := range testCases {
			code := RunCmd(testCase.cmd, Environment{})
			assert.Equal(t, testCase.expectedCode, code)
		}
	})
}
