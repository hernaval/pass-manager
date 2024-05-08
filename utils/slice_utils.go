package utils

func Filter[T any](slice []T, test func(T) bool) (ret []T) {
	for _, s := range slice {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
