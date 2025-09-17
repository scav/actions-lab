package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Me(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		result := 2 + 2
		require.Equal(t, 4, result)
	})

	t.Run("multi", func(t *testing.T) {
		result := 2 * 2
		require.Equal(t, 4, result)
	})

	t.Run("err", func(t *testing.T) {
		result := 2 / 2
		require.NotEqual(t, 4, result)
	})

	println("test dump")
}
