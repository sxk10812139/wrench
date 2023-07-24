package wrench

import "strconv"

// StringToUint64 String2Uint64
func StringToUint64(str string) uint64 {
	i64, _ := strconv.Atoi(str)
	u64 := uint64(i64)

	return u64
}

// StringToInt64 StringToInt64
func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

// StringToInt String2Int
func StringToInt(str string) int {
	intNum, _ := strconv.Atoi(str)
	return intNum
}

// Int642String Int642String
func Int64ToString(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

// Uint642String
func Uint64ToString(i uint64) string {
	str := strconv.FormatUint(i, 10)
	return str
}

// Int2String Int2String
func IntToString(i int) string {
	str := strconv.Itoa(i)
	return str
}

// Float64ToString
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

func SliceToMap[T any, K comparable, V any](s []T, f func(e T) (K, V)) map[K]V {
	o := make(map[K]V)
	for _, e := range s {
		key, value := f(e)
		o[key] = value
	}
	return o
}

func MapToSlice[T any, K comparable, V any](in map[K]V, f func(k K, v V) T) []T {
	o := make([]T, 0)
	for k, v := range in {
		e := f(k, v)
		o = append(o, e)
	}
	return o
}
