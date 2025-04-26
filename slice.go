package slicezh

import (
	"math/rand/v2"
	"slices"
)

func Do过滤[T any](a []T, condition func(value T) bool) []T {
	results := make([]T, 0, len(a))
	for _, one := range a {
		if condition(one) {
			results = append(results, one)
		}
	}
	return results
}

func Do翻转[T any](a []T) {
	slices.Reverse(a)
}

func Get逆序新数组[T any](a []T) []T {
	results := make([]T, len(a))
	for idx, one := range a {
		results[len(results)-idx-1] = one
	}
	return results
}

func Get随机取样[T any](a []T) T {
	return a[rand.IntN(len(a))]
}
