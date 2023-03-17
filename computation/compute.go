package computation

import (
	"errors"
)

// Compute returns an array of successive integers starting by 1
// incremented by 1 and of size `limit`,
// where all multiples of `int1` are replaced by `str1`,
// all multiples of `int2` are replaced by `str2`,
// all multiples of `int1` and `int2` are replaced by concatenation of `str1` and `str2`.
func Compute(int1, int2, limit uint64, str1, str2 string) ([]interface{}, []error) {

	errs := make([]error, 0)

	if int1 == 0 {
		errs = append(errs, errors.New("int1 must be greater than 0"))
	}

	if int2 == 0 {
		errs = append(errs, errors.New("int2 must be greater than 0"))
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
	result := make([]interface{}, limit)

	if limit == 0 {
		// Useless to go deeper
		return result, errs
	}

	test_int1 := int1 <= limit
	test_int2 := int2 <= limit

	// TODO:
	//	If limit will be change to an *uint64,
	// 	then need to find a solution about how to manage
	//	when int1*int2 > ^uint64(0)
	test_int1_multiply_by_int2 := (limit / int1) >= int2

	int1_multiply_by_int2 := uint64(0)
	if test_int1_multiply_by_int2 {
		// we know limit is an uint64 so we don't care about result of int1*int2
		int1_multiply_by_int2 = int1 * int2
	}

	str1str2 := str1 + str2
	current_value := uint64(1)

	// Note: 
	//	Idea for improvements: Split result and use goroutines to fill each part

	for a := range result {
		if test_int1_multiply_by_int2 && current_value%int1_multiply_by_int2 == 0 {
			result[a] = str1str2
		} else {
			if test_int1 && current_value%int1 == 0 {
				result[a] = str1
			} else if test_int2 && current_value%int2 == 0 {
				result[a] = str2
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
