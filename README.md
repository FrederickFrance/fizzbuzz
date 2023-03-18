# FIZZBUZZ

Web server that exposes two REST API endpoints that:

- Accepts five parameters: three integers `int1`, `int2` and `limit`, and two strings `str1` and `str2`. Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
- Returns statistics allowing users to know what the most frequent request has been. Returns the parameters corresponding to the most used request, as well as the number of hits for this request.

## How to launch

```
PORT=8080 go run .
```

## How to use REST API endpoint

Call http://localhost:8080/compute?int1=3&int2=5&limit=20&str1=fizz&str2=buzz to return a list of string corresponding to the statement of the fizzbuzz exercise.

Call http://localhost:8080/metrics to return a list of string corresponding to the statement of the fizzbuzz exercise.

## Limitations

- `int1` and `int2` must be strictly positive.
- `limit` must be positive.
- `int1`, `int2` and `limit` cannot be higher than 18446744073709551615 (maximum value for uint64 type).
- Need to add a linter
- Need to launch tests from github
- Need to add an OpenApi documentation
- `limit` cannot be higher than 1000, `str1` and `str2` cannot be longer than 10 characters. It's an arbitrary hard-coded limit to avoid crash with an out-of-memory exception.
- The counter for the statistics endpoint saves only parameters when int1, int2 and limit are uint64 and str1 and str2 are strings, because of lack of specifications.
- The statistics endpoint could be replaced with [Prometheus](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus).
- The statistics endpoint returns only one element (in case of equal values, the original order is used).


## Implementation decision

Instead of returning a list of strings, the function returns an array of interface to avoid conversation from number to string.
