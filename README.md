# FIZZBUZZ

Web server that exposes a REST API endpoint that:

- Accepts five parameters: three integers `int1`, `int2` and `limit`, and two strings `str1` and `str2`.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

## How to launch

```
PORT=8080 go run .
```

## How to use REST API endpoint

Call http://localhost:8080/compute?int1=3&int2=5&limit=20&str1=fizz&str2=buzz

## Limitations

- `int1` and `int2` must be strictly positive.
- `limit` must be positive.
- `int1`, `int2` and `limit` cannot be higher than 18446744073709551615 (maximum value for uint64 type).
- Need to add a linter
- Need to launch tests from github
- Need to add an OpenApi documentation
- Currently, can crash with an out-of-memory exception when `limit` is huge. Need to fix if necessary. 

## Implementation decision

Instead of returning a list of strings, the function returns an array of interface to avoid conversation from number to string.
