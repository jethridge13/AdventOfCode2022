package util

func SlicePop[T any](s []T, i int) (T, []T) {
	n := s[i]
	s = append(s[:i], s[i+1:]...)
	return n, s
}
