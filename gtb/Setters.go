package gtb

import (
	"errors"
	"fmt"
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
		EoE(fmt.Errorf("Error Accessing relPath, %w", os.Chdir(path)))
	} else {
		EoE(fmt.Errorf("Error Getting Caller Location: %w", errors.New(filename)))
	}
}
