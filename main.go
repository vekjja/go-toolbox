package gtills

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

// E : exit with error code 1 and print if err is notnull
func E(err error, msg string) {
	if err != nil {
		fmt.Println("❌  "+msg, err)
		os.Exit(1)
	}
}

func GetHomeDir() string {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	} else {
		// Maybe it's cross compilation without cgo support. (darwin, unix)
		return os.Getenv("HOME")
	}
}

func GetInput(q string) string {
	if q != "" {
		fmt.Print(q)
	}
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

func SetFromInput(a *string, q string) {
	*a = strings.TrimRight(GetInput(q), "\n")
}
