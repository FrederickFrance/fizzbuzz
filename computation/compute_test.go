package computation_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"test.com/fizzbuzz/computation"
)

func TestComputation(t *testing.T) {
	values, errs := computation.Compute(3, 5, 20, "fizz", "buzz")

	assert.Empty(t, errs)

	assert.Equal(t,
		[]interface{}{
			uint64(1), uint64(2), "fizz", uint64(4), "buzz",
			"fizz", uint64(7), uint64(8), "fizz", "buzz",
			uint64(11), "fizz", uint64(13), uint64(14), "fizzbuzz",
			uint64(16), uint64(17), "fizz", uint64(19), "buzz",
		},
		values,
	)

	values, errs = computation.Compute(0, 5, 20, "fizz", "buzz")

	assert.Empty(t, values)

	assert.Equal(t,
		[]error{
			errors.New("int1 must be greater than 0"),
		},
		errs,
	)

	values, errs = computation.Compute(0, 0, 20, "fizz", "buzz")

	assert.Empty(t, values)

	assert.Equal(t,
		[]error{
			errors.New("int1 must be greater than 0"),
			errors.New("int2 must be greater than 0"),
		},
		errs,
	)

	values, errs = computation.Compute(123, 456, 0, "fizz", "buzz")

	assert.Empty(t, values)

	assert.Empty(t, errs)

	values, errs = computation.Compute(3, 5, 2, "fizz", "buzz")

	assert.Empty(t, errs)

	assert.Equal(t,
		[]interface{}{uint64(1), uint64(2)},
		values,
	)

	values, errs = computation.Compute(3, 1, 7, "fizz", "buzz")

	assert.Empty(t, errs)

	assert.Equal(t,
		[]interface{}{
			"buzz", "buzz", "fizzbuzz", "buzz", "buzz",
			"fizzbuzz", "buzz",
		},
		values,
	)
}
