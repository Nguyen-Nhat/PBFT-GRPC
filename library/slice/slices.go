package slice

const (
	ZERO_INT = 0
)

func Make[T any](values ...T) []T {
	res := make([]T, 0, len(values))
	if len(values) > ZERO_INT {
		res = append(res, values...)
	}
	return res
}
