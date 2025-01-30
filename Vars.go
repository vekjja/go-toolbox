package toolbox

import (
	"encoding/gob"
	"os"
)

// WriteVar : write gob to local filesystem
func WriteVar(file string, data interface{}) error {
	gobFile, err := os.Create(file)
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(gobFile)
	encoder.Encode(data)
	gobFile.Close()
	return nil
}

// ReadVar : read gob from local filesystem
func ReadVar(file string, object interface{}) error {
	gobFile, err := os.Open(file)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(gobFile)
	err = decoder.Decode(object)
	gobFile.Close()
	return nil
}
