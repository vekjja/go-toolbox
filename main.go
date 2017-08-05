package gtils

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

// EoE : exit with error code 1 and print if err is notnull
func EoE(err error, msg string) {
	if err != nil {
		fmt.Println("❌  "+msg, err)
		os.Exit(1)
	}
}

// GetHomeDir : returns a full path to user's home dorectory
func GetHomeDir() string {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	}
	// Maybe it's cross compilation without cgo support. (darwin, unix)
	return os.Getenv("HOME")
}

// Confirm : return confirmation based on user input
func Confirm(q string) bool {
	a := GetInput(q + " (Y/n) ")
	var res bool
	switch a {
	case "":
		fallthrough
	case "y":
		fallthrough
	case "Y":
		res = true
	case "n":
	case "N":
		res = false
	default:
		return Confirm(q)
	}
	return res
}

// GetInput : return string of user input
func GetInput(q string) string {
	if q != "" {
		fmt.Print(q)
	}
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

// SetFromInput : set value of provided var to the value of user input
func SetFromInput(a *string, q string) {
	*a = strings.TrimRight(GetInput(q), "\n")
}
