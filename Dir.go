package toolbox

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// Mkdir : make a directory if it does not exist
func Mkdir(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}
}

// SetDir : resolves the absolute path for provided relative path and sets the working directory.
func SetDir(relPath string) error {
	if _, filename, _, ok := runtime.Caller(1); ok {
		re := regexp.MustCompile("[a-zA-Z0-9-]*.go$")
		path := filepath.Join(re.ReplaceAllString(filename, ""), relPath)
		return os.Chdir(path)
	} else {
		return fmt.Errorf("Error Getting Caller Location: %s", filename)
	}
}
