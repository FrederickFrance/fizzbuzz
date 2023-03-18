package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func MostUsed[K comparable, V Number](m map[K]V) *K {
	// XXX: If necessary, replace this code by using an heap

	if len(m) == 0 {
		return nil
	}

	keys := make([]K, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return &keys[0]
}
