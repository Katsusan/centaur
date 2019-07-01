package util

import (
	"os"
)

//FileExist will return true if file exists and can be accessed,
//otherwise returns false
func FileExist(file string) bool {
	fileinfo, err := os.Stat(file)
	return err == nil && !fileinfo.IsDir()
}
