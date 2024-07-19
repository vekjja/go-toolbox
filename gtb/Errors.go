package gtb

import (
	"errors"
	"fmt"
	"os"
)

func formatError(err error, msg ...string) error {
	if len(msg) > 0 {
		errMsg := ""
		for i, m := range msg {
			if i > 0 {
				errMsg += "   " + m + "\n"
			} else {
				errMsg += Red(m) + "\n"
			}
		}
		err = errors.New(errMsg)
	} else {
		err = errors.New(Red(err.Error()))
	}
	return fmt.Errorf("\n💔 %s", err)
}

// EoE : if err is not nil format and print error then exit
func EoE(err error, msg ...string) {
	if err != nil {
		fmt.Println(formatError(err, msg...))
		os.Exit(1)
	}
}
