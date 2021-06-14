package utils

import (
	"io/fs"
	"runtime"
)

// SourceFilePath ...
func SourceFilePath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fs.ErrInvalid
	}
	return file, nil
}
