package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		dir, _ := os.MkdirTemp("./testdata", "env2")
		defer os.RemoveAll(dir)
		env, err := ReadDir(dir)
		assert.NoError(t, err)
		assert.Len(t, env, 0)
	})

	t.Run("empty dir", func(t *testing.T) {
		dir, _ := os.MkdirTemp("./testdata/", "env3")
		defer os.RemoveAll(dir)
		os.CreateTemp(dir, "=SKIP")
		os.MkdirTemp(dir, "SKIPDIR")
		os.CreateTemp(dir, "EMPTY1")
		f4, _ := os.CreateTemp(dir, "FILLED")
		f4.WriteString("   space    ")
		f4Info, _ := f4.Stat()

		env, err := ReadDir(dir)
		assert.NoError(t, err)
		assert.Len(t, env, 2)
		assert.Equal(t, "   space", env[f4Info.Name()].Value)
	})

	t.Run("incorrect dir", func(t *testing.T) {
		env, err := ReadDir("./testdata/fakedir3213")
		require.Truef(t, errors.Is(err, ErrIncorrectDirectory), "actual err - %v", err)
		assert.Len(t, env, 0)
	})
}
