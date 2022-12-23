package util

import "fmt"

func PrintGraph[T any](graph [][]T) {
	for _, row := range graph {
		fmt.Println(row)
	}
}
