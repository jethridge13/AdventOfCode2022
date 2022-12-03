package util

func FindMax(s []int) int {
	max := 0
	for _, i := range s {
		if i > max {
			max = i
		}
	}
	return max
}
