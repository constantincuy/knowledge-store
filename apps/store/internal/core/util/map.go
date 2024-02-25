package util

func Map[T any, R any](values []T, mapper func(T) R) []R {
	result := make([]R, len(values))
	for i, v := range values {
		result[i] = mapper(v)
	}

	return result
}
