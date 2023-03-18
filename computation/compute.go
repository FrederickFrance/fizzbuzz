package computation

import (
	"errors"

	"test.com/fizzbuzz/metrics"
)

const Uint64Limit = uint64(1_000)
const StringLimit = 10

// Compute returns an array of successive integers starting by 1
// incremented by 1 and of size `limit`,
// where all multiples of `int1` are replaced by `str1`,
// all multiples of `int2` are replaced by `str2`,
// all multiples of `int1` and `int2` are replaced by concatenation of `str1` and `str2`.
func Compute(parameters metrics.Parameters) ([]interface{}, []error) {

	errs := make([]error, 0)

	if parameters.Int1 == 0 {
		errs = append(errs, errors.New("int1 must be greater than 0"))
	}

	if parameters.Int2 == 0 {
		errs = append(errs, errors.New("int2 must be greater than 0"))
	}

	// TODO:
	// Need to improve.
	// Arbitrary limits to avoid out-of-memory crash.
	if parameters.Int1 > Uint64Limit || parameters.Int2 > Uint64Limit || parameters.Limit > Uint64Limit ||
		len(parameters.Str1) > StringLimit || len(parameters.Str2) > StringLimit {
		errs = append(errs, errors.New("int1, int 2 and limit must not be higher than 1_000. str1 and str2 must not be longer than 10 characters."))
	}

	if len(errs) > 0 {
		return make([]interface{}, 0), errs
	}

	/*
		 TODO:
		 	Improve when golang will be able to init an array
			with a begin and an increment.
			1. Then loop with int1 as increment
			2. Then loop with int2 as increment
			3. Then loop with int1*int2 as increment
	*/
	result := make([]interface{}, parameters.Limit)

	if parameters.Limit == 0 {
		// Useless to go deeper
		return result, errs
	}

	test_int1 := parameters.Int1 <= parameters.Limit
	test_int2 := parameters.Int2 <= parameters.Limit

	// TODO:
	//	If limit will be change to an *uint64,
	// 	then need to find a solution about how to manage
	//	when int1*int2 > ^uint64(0)
	test_int1_multiply_by_int2 := (parameters.Limit / parameters.Int1) >= parameters.Int2

	int1_multiply_by_int2 := uint64(0)
	if test_int1_multiply_by_int2 {
		// we know limit is an uint64 so we don't care about result of int1*int2
		int1_multiply_by_int2 = parameters.Int1 * parameters.Int2
	}

	str1str2 := parameters.Str1 + parameters.Str2
	current_value := uint64(1)

	// Note:
	//	Idea for improvements: Split result and use goroutines to fill each part

	for a := range result {
		if test_int1_multiply_by_int2 && current_value%int1_multiply_by_int2 == 0 {
			result[a] = str1str2
		} else {
			if test_int1 && current_value%parameters.Int1 == 0 {
				result[a] = parameters.Str1
			} else if test_int2 && current_value%parameters.Int2 == 0 {
				result[a] = parameters.Str2
			} else {
				// XXX:
				//	If result must be an array of strings,
				//	then need to use strconv.FormatUint(uint64(current_value), 10)
				result[a] = current_value
			}
		}

		current_value += 1
	}

	return result, errs
}
