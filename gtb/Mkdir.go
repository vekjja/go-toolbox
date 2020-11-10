package gtb

import "os"

// Mkdir : make a directory if it does not exist
func Mkdir(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}
}
