package utils_test

import (
	"testing"

	"github.com/go-zwbc/slicezh/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestZero(t *testing.T) {
	require.Equal(t, 0, utils.Zero[int]())
	require.Equal(t, "", utils.Zero[string]())
	require.Equal(t, 0.0, utils.Zero[float64]())
	require.Equal(t, rune(0), utils.Zero[rune]())
}
