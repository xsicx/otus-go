package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("unsupported file error", func(t *testing.T) {
		fromPath := "broken111File"
		toPath := "out.txt"
		var offset, limit int64
		err := Copy(fromPath, toPath, offset, limit)

		require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual err - %v", err)
	})

	t.Run("offset error", func(t *testing.T) {
		fromPath := "testdata/input2.txt" // 2 bytes
		toPath := "out.txt"
		var offset int64 = 100
		var limit int64
		err := Copy(fromPath, toPath, offset, limit)

		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual err - %v", err)
	})
}
