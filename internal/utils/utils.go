package utils

import (
	"io"
	"os"
)

func OpenFile(fileName string) io.Reader {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func Map[T any, R any](slice []T, mapFn func(T) R) []R {
	result := make([]R, 0)
	for idx := 0; idx < len(slice); idx++ {
		result = append(result, mapFn(slice[idx]))
	}
	return result
}

func Filer[T any](slice []T, filerFn func(T) bool) []T {
	result := make([]T, 0)
	for idx := 0; idx < len(slice); idx++ {
		if filerFn(slice[idx]) {
			result = append(result, slice[idx])
		}
	}
	return result
}
