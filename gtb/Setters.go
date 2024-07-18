package gtb

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// SetRelPath : resolves the absolute path for provided relative path.
func SetRelPath(relPath string) {
	if _, filename, _, ok := runtime.Caller(1); ok {
		re := regexp.MustCompile("[a-zA-Z0-9-]*.go$")
		path := filepath.Join(re.ReplaceAllString(filename, ""), relPath)
		EoE(os.Chdir(path), "Error Accessing relPath:")
	} else {
		EoE(errors.New(filename), "Error Getting Caller Location")
	}
}
