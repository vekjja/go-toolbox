package toolbox

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenBrowser : open platform specific windows
func OpenBrowser(url string) (*exec.Cmd, error) {
	var err error
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		err = fmt.Errorf("Unsupported Platform")
	}

	return cmd, err

}
