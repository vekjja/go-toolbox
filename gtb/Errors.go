package gtb

import (
	"errors"
	"fmt"
	"os"
)

func FormatError(err error, msg ...string) error {
	if err == nil {
		return nil
	}
	errMsg := ""
	if len(msg) > 0 {
		for _, m := range msg {
			errMsg += Red(m)
		}
	}
	errMsg += "\n" + Red(err.Error())
	return fmt.Errorf("\n💔 %s", Red(errors.New(errMsg)))
}

func EoE(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
