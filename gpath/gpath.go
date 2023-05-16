package gpath

import (
	"os"
)

func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
