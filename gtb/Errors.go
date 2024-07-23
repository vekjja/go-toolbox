package gtb

import (
	"fmt"
	"os"
)

func formatError(err error, msg ...string) error {
	if len(msg) > 0 {
		errMsg := err.Error()
		for _, m := range msg {
			errMsg += m
		}
		err = fmt.Errorf(errMsg)
	}
	return fmt.Errorf("\n💔 %s", Red(err.Error()))
}

// EoE : if err is not nil format and print error then exit
func EoE(err error, msg ...string) {
	if err != nil {
		fmt.Println(formatError(err, msg...))
		os.Exit(1)
	}
}
