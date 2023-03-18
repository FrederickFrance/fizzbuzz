package metrics

type Parameters struct {
	Int1  uint64
	Int2  uint64
	Limit uint64
	Str1  string
	Str2  string
}

// Map to save how many times each set of parameters has been called.
//
// TODO:
//
// - Need to add a database to keep data.
//
// - Need to use sync.Mutex if UsedParameters becomes a shared map.
var UsedParameters = make(map[Parameters]uint64)
