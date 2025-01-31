package toolbox

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

// ExecPath :
func ExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// FilesInDir : return an array of os.Fileinfo for the given path
func FilesInDir(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// HomeDir : returns a full path to user's home directory
func HomeDir() string {
	home, err := os.UserHomeDir()
	EoE(err)
	return home
}

// Input : return string of user input
func Input() string {
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

// UserInfo :
func UserInfo(userName string, uID int) *user.User {

	var usr *user.User
	var err error

	if userName != "" {
		usr, err = user.Lookup(userName)
	} else {
		usr, err = user.LookupId(strconv.Itoa(uID))
	}
	if err != nil {
		log.Fatal(err)
	}
	return usr
}
