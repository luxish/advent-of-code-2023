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
