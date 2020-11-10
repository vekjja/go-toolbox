package gtb

import (
	"fmt"
	"os"
	"testing"
)

// TestGetUserInfo : test the SplitMulti func
func TestOpenBrowser(t *testing.T) {

	cmd, err := OpenBrowser("https://google.com")
	if err != nil {
		t.Errorf("Error Opening Browser: " + err.Error())
	}

	fmt.Println(cmd)

	// err = cmd.Start()
	// if err != nil {
	// 	t.Errorf("Error Opening Browser " + err.Error())
	// }

	// cmd.Process.Kill()

}

// TestGetUserInfo : test the SplitMulti func
func TestGetUserInfo(t *testing.T) {

	userName := "root"

	usr := GetUserInfo(userName, 0)

	if AToUint32(usr.Uid) != uint32(0) {
		t.Errorf("Expected UID 0, got %v", usr.Uid)
	}

	userName = ""
	usr = GetUserInfo(userName, 0)
	if usr.Username != "root" {
		t.Errorf("Expected User root, got %v", usr.Username)
	}

}

// TestSplitMulti : test the SplitMulti func
func TestSplitMulti(t *testing.T) {

	testString := "foo bar baz . oof rab zab"

	split := SplitMulti(testString, " .")

	if len(split) != 6 {
		t.Errorf("Number of words split = %d; want 6", len(split))
	}

	testString = "foo bar [baz]"

	split = SplitMulti(testString, " []")

	if len(split) != 3 {
		t.Errorf("Number of words split = %d; want 3", len(split))
	}

}

// TestAToUint32 : test the AToUint32 func
func TestAToUint32(t *testing.T) {

	testString := "42"

	testUint := AToUint32(testString)

	if uint32(42) != testUint {
		t.Errorf("Converted String = %d; want uint(42)", testUint)
	}

}

// TestGetFilesInDir : test the GetFilesInDir func
func TestGetFilesInDir(t *testing.T) {

	testDir := "./test"

	defer os.RemoveAll(testDir)
	err := os.Mkdir(testDir, 0755)

	if err != nil {
		t.Errorf("Error making test directory: %v", err)
	}

	testFiles := []string{"one", "two", "three"}
	for _, file := range testFiles {
		os.Create(fmt.Sprintf("%v/%v", testDir, file))
	}

	files := GetFilesInDir(testDir)

	if len(files) != 3 {
		t.Errorf("Expected 3 files, got %d", len(files))
	}

}
