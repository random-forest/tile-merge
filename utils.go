package main

import "os"

func Min(a ...int) int {
	min := int(^uint(0) >> 1)
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MakeDirs(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return
		}
	}
}

func Range(start, end, step int) []int {
	var out []int

	for i := start; i <= end; i += step {
		out = append(out, i)
	}

	return out
}

func FRange(start, end, step float64) []float64 {
	var out []float64

	for i := start; i <= end; i += step {
		out = append(out, i)
	}

	return out
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
