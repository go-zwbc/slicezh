package slicezh

import (
	"math/rand/v2"
	"slices"

	"github.com/go-zwbc/slicezh/internal/utils"
)

func Get按条件过滤[T any](a []T, condition func(value T) bool) []T {
	results := make([]T, 0, len(a))
	for _, one := range a {
		if condition(one) {
			results = append(results, one)
		}
	}
	return results
}

func Get首个满足条件的[T any](a []T, condition func(value T) bool) (T, bool) {
	for _, one := range a {
		if condition(one) {
			return one, true
		}
	}
	return utils.Zero[T](), false
}

func Do翻转自身内容[T any](a []T) {
	slices.Reverse(a)
}

func Get获取逆序新数组[T any](a []T) []T {
	results := make([]T, len(a))
	for idx, one := range a {
		results[len(results)-idx-1] = one
	}
	return results
}

func Get随机取个元素[T any](a []T) T {
	return a[rand.IntN(len(a))]
}

func Contains[T comparable](a []T, value T) bool {
	return slices.Contains(a, value)
}

func In[T comparable](value T, a []T) bool {
	return slices.Contains(a, value)
}

func Get转得Map[K comparable, V any](a []V, keyFunc func(value V) K) map[K]V {
	resMap := make(map[K]V, len(a))
	for _, one := range a {
		resMap[keyFunc(one)] = one
	}
	return resMap
}
