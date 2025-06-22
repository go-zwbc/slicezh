package slicezh_test

import (
	"testing"

	"github.com/go-zwbc/slicezh"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestGet按条件过滤(t *testing.T) {
	results := slicezh.Get按条件过滤([]int{1, 2, 3, 4, 5}, func(value int) bool {
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

func TestMut翻转自身内容(t *testing.T) {
	a := []string{"a", "b", "c"}
	slicezh.Mut翻转自身内容(a)
	require.Equal(t, []string{"c", "b", "a"}, a)
}

func TestGet获取逆序新数组(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := slicezh.Get获取逆序新数组(a)
	require.Equal(t, []string{"c", "b", "a"}, b)
	require.Equal(t, []string{"a", "b", "c"}, a)
}

func TestGet随机取个元素(t *testing.T) {
	v := slicezh.Get随机取个元素([]float64{1.0, 2.14, 3.1})
	t.Log(v)
}

func TestMut随机乱序元素(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f", "g"}
	slicezh.Mut随机乱序元素(a)
	t.Log(a)
}

func TestContains(t *testing.T) {
	require.True(t, slicezh.Contains([]int{1, 2, 3}, 2))
	require.False(t, slicezh.Contains([]int{1, 2, 3}, 4))
}

func TestIn(t *testing.T) {
	require.True(t, slicezh.In(2, []int{1, 2, 3}))
	require.False(t, slicezh.In(0, []int{1, 2, 3}))
}

func TestGet映射Map(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []*Person{
		{Name: "杨亦乐", Age: 18},
		{Name: "刘亦菲", Age: 18},
		{Name: "古天乐", Age: 18},
	}
	resMap := slicezh.Get映射Map(people, func(value *Person) string {
		return value.Name
	})
	require.Equal(t, 3, len(resMap))
	t.Log(neatjsons.S(resMap))
}
