package gtb

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

// GetExecPath :
func GetExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// GetFilesInDir : return an array of os.Fileinfo for the given path
func GetFilesInDir(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// GetHomeDir : returns a full path to user's home directory
func GetHomeDir() string {
	home, err := os.UserHomeDir()
	EoE(err)
	return home
}

// GetInput : return string of user input
func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	return strings.TrimRight(ans, "\n")
}

// GetUserInfo :
func GetUserInfo(userName string, uID int) *user.User {

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
