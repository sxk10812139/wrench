package wrench

func Slice2Map[T any, K comparable, V any](s []T, f func(e T) (K, V)) map[K]V {
	o := make(map[K]V)
	for _, e := range s {
		key, value := f(e)
		o[key] = value
	}
	return o
}

func Map2Slice[T any, K comparable, V any](in map[K]V, f func(k K, v V) T) []T {
	o := make([]T, 0)
	for k, v := range in {
		e := f(k, v)
		o = append(o, e)
	}
	return o
}
