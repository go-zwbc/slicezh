package slicezh_test

import (
	"testing"

	"github.com/go-zwbc/slicezh"
	"github.com/stretchr/testify/require"
)

func TestGet按过滤过滤(t *testing.T) {
	results := slicezh.Get按过滤过滤([]int{1, 2, 3, 4, 5}, func(value int) bool {
		return value%2 == 0
	})
	t.Log(results)
	require.Equal(t, []int{2, 4}, results)
}

func TestGet首个满足条件的(t *testing.T) {
	one, ok := slicezh.Get首个满足条件的([]int{1, 2, 3, 4, 5}, func(value int) bool {
		return value%2 == 0
	})
	require.True(t, ok)
	t.Log(one)
	require.Equal(t, 2, one)
}

func TestDo翻转(t *testing.T) {
	a := []string{"a", "b", "c"}
	slicezh.Do翻转(a)
	require.Equal(t, []string{"c", "b", "a"}, a)
}

func TestGet逆序新数组(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := slicezh.Get逆序新数组(a)
	require.Equal(t, []string{"c", "b", "a"}, b)
	require.Equal(t, []string{"a", "b", "c"}, a)
}

func TestGet随机取样(t *testing.T) {
	v := slicezh.Get随机取样([]float64{1.0, 2.14, 3.1})
	t.Log(v)
}

func TestContains(t *testing.T) {
	require.True(t, slicezh.Contains([]int{1, 2, 3}, 2))
	require.False(t, slicezh.Contains([]int{1, 2, 3}, 4))
}

func TestIn(t *testing.T) {
	require.True(t, slicezh.In(2, []int{1, 2, 3}))
	require.False(t, slicezh.In(0, []int{1, 2, 3}))
}
