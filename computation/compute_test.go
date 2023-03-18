package computation_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"test.com/fizzbuzz/computation"
	"test.com/fizzbuzz/metrics"
)

func compute(int1, int2, limit uint64, str1, str2 string) ([]interface{}, []error) {
	return computation.Compute(
		metrics.Parameters{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
		})
}

func TestComputation(t *testing.T) {
	values, errs := compute(3, 5, 20, "fizz", "buzz")

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

	values, errs = compute(0, 5, 20, "fizz", "buzz")

	assert.Empty(t, values)

	assert.Equal(t,
		[]error{
			errors.New("int1 must be greater than 0"),
		},
		errs,
	)

	values, errs = compute(0, 0, 20, "fizz", "buzz")

	assert.Empty(t, values)

	assert.Equal(t,
		[]error{
			errors.New("int1 must be greater than 0"),
			errors.New("int2 must be greater than 0"),
		},
		errs,
	)

	values, errs = compute(123, 456, 0, "fizz", "buzz")

	assert.Empty(t, values)

	assert.Empty(t, errs)

	values, errs = compute(3, 5, 2, "fizz", "buzz")

	assert.Empty(t, errs)

	assert.Equal(t,
		[]interface{}{uint64(1), uint64(2)},
		values,
	)

	values, errs = compute(3, 1, 7, "fizz", "buzz")

	assert.Empty(t, errs)

	assert.Equal(t,
		[]interface{}{
			"buzz", "buzz", "fizzbuzz", "buzz", "buzz",
			"fizzbuzz", "buzz",
		},
		values,
	)

	values, errs = compute(1, 1, 1000, "1234567890", "1234567890")
	assert.Empty(t, errs)
	assert.NotEmpty(t, values)

	values, errs = compute(1, 1, 1001, "1234567890", "1234567890")
	assert.NotEmpty(t, errs)
	assert.Empty(t, values)

	values, errs = compute(1001, 1, 1001, "1234567890", "1234567890")
	assert.NotEmpty(t, errs)
	assert.Empty(t, values)

	values, errs = compute(1001, 1, 1, "1234567890", "1234567890")
	assert.NotEmpty(t, errs)
	assert.Empty(t, values)

	values, errs = compute(1, 1001, 1, "1234567890", "1234567890")
	assert.NotEmpty(t, errs)
	assert.Empty(t, values)

	values, errs = compute(1, 1, 1, "1", "12345678901")
	assert.NotEmpty(t, errs)
	assert.Empty(t, values)

	values, errs = compute(1, 1, 1, "12345678901", "1")
	assert.NotEmpty(t, errs)
	assert.Empty(t, values)

}
