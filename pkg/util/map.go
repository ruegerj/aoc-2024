package util

import (
	"cmp"
	"errors"
	"slices"
)

func KeyByValue[K comparable, V comparable](m map[K]V, value V) K {
	for k, v := range m {
		if v == value {
			return k
		}
	}

	panic(errors.New("Map doesn't contain a key which maps to the provided value"))
}

func SortedKeys[K cmp.Ordered, V any](m map[K]V, desc bool) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	slices.Sort(keys)

	if desc {
		slices.Reverse(keys)
	}

	return keys
}
