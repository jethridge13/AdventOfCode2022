package util

func GetManhattanDistance(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}
