package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"test.com/fizzbuzz/utils"
)

func TestMostUsed(t *testing.T) {

	m := map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	assert.Equal(t, "bacon", *utils.MostUsed(m))
}
