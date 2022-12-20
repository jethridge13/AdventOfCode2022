package util

func SliceInsert[T any](seq []T, i int, item T) []T {
	seq = append(seq[:i+1], seq[i:]...)
	seq[i] = item
	return seq
}
